### Диаграмма

```plantuml
sequenceDiagram
    autonumber
    actor C as Client
    participant H  as Handler<br/>v5/fullstat
    participant DR as decoder.go
    participant DD as DayDepthChecker
    participant S  as stat.Service
    participant AF as adverts.Functionality
    participant PA as PostgreSQL<br/>"Adverts"
    participant G  as awesomestat<br/>gRPC (external)
    participant SS as SupplierSettings
    participant R  as Redis (cache)
    participant PS as PostgreSQL<br/>Suppliers / settings
    participant ST as Stocker
    participant PC as PostgreSQL<br/>contents_full
    participant TF as catalogs/search/<br/>recom/auto/seacat
    participant PN as PostgreSQL<br/>NM tables

    C->>H: GET /v5/fullstat?advertID,supplierID,from,to,...
    H->>DR: decodeRequest(r)
    DR-->>H: FullStatRequestParams
    H->>DD: DayDepthChecker(begin,end)
    DD-->>H: ok / 400
    H->>S: GetFullStatsV5(ctx, req)

    %% --- ownership ---
    S->>AF: CheckSupplier(supplierID, advertID)
    AF->>PA: SELECT exists("Adverts" by Id+SupplierId)
    PA-->>AF: bool
    AF-->>S: own?
    Note over S: !own → ErrWrongSupplier (400)

    %% --- advert info ---
    S->>AF: AdvertShortInfo(advertID)
    AF->>PA: SELECT Type,StatusId,CampaignName,CreatedAt,EndDate
    PA-->>AF: row
    AF-->>S: AdvertWithCampaignShortInfo
    Note over S: status ∉ {7,9,11} → ErrDeletedOrRejected<br/>IsDateRangeOutOfCampaignPeriod → ErrNoContent (204)<br/>UpdateCampaignDates(сужение диапазона)<br/>AdvertTypeFromInt(advert.Type)

    %% --- main stats from gRPC ---
    alt PlacementType == All
        S->>G: GetStatsDailyAndNm(supp,adv,from,to)
    else иначе
        S->>G: GetStatsDailyAndNmByPlacement(...,placement)
    end
    G-->>S: model.FullStat<br/>(DailyStats, NMStats, ImtNmStats, Previous, Stat)

    %% --- supplier settings (1) ---
    S->>SS: SupplierSettings(supplierID)
    SS->>R: GET supplier:{id}:settings
    alt cache hit
        R-->>SS: JSON
    else cache miss
        R-->>SS: ErrNotFound
        SS->>PS: SELECT Suppliers (CountryCode,...)
        PS-->>SS: SupplierInfo
        SS->>PS: SELECT global_ads_settings ∪ suppliers_ads_settings
        PS-->>SS: rawSettings
        SS->>R: SET supplier:{id}:settings (TTL 30m)
    end
    SS-->>S: map[string]string

    %% --- multicard branch ---
    alt withMulticard == true (ui_multicard_statistics_details_enabled)
        Note over S: filterMulticardNms (убираем дубли nm,<br/>оставляем заголовок imt и его детей)<br/>filterMulticardMetrics (Views==0 ⇒ обнуление)<br/>сортировка imtNmStats by views DESC
    else
        Note over S: removeMulticardStats<br/>(удаляем все строки с ImtNmStats)
    end

    Note over S: filterFullstatMetrics(clearDrrMetrics)<br/>обнуляет Drr в Stat / DailyStats / NMStats / ImtNmStats

    alt len(DailyStats) == 0
        S-->>H: stats.ErrNoStats
        H-->>C: 202 «Статистика в процессе получения»
    end

    Note over S: sort DailyStats by Date<br/>MergeDailyStatsSameDate<br/>AppendEmptyDailyStat<br/>set Begin / End / AdvertID

    %% --- enrich names ---
    S->>ST: EnrichNMs(all NMIDs incl. imt)
    ST->>PC: SELECT contents_full × kinds<br/>LEFT JOIN discount_nms WHERE nm_id = any($1)
    PC-->>ST: NmKindSubject[]
    ST-->>S: products
    Note over S: проставляет Name в NMStats и в ImtNmStats

    Note over S: sortNMStats(order, direction)

    %% --- ad-type NMs (фактически не используется) ---
    S->>TF: AdvertNMs(advertID) (по типу кампании)
    alt Catalog
        TF->>PN: SELECT array_agg FROM catalog_ads
    else Search
        TF->>PN: JOIN item_costs × advert_nomenclatures × filters
    else Recom
        TF->>PN: SELECT FROM recom_ads
    else Auto
        TF->>PN: SELECT FROM advert_nomenclatures (active)
    else SearchCatalog (seacat)
        TF->>PN: тот же запрос, что у Search
    end
    PN-->>TF: []nmIDs
    TF-->>S: []nmIDs
    Note over S: результат не используется<br/>(потенциальный мёртвый код)

    %% --- split / glue ---
    Note over S: separateAdvertNMStats:<br/>"nm","imt" → advertNMStats<br/>"supplier" → SideNMStats
    alt request.GlueNmStats == true
        Note over S: nmStats = glueNmStatsPerNmID(всех NMStats)<br/>(агрегация sum по NMID, AttributionStage игнорируется)
    else
        Note over S: nmStats = advertNMStats
    end

    Note over S: fullStat.NMStats = nmStats<br/>fullStat.Stat = recalcTotalFromShownNMStats(...)<br/>fullStat.AppType = req.AppType

    %% --- supplier settings (2) ---
    S->>SS: SupplierSettings(supplierID)
    SS->>R: GET supplier:{id}:settings
    R-->>SS: JSON (обычно hit)
    SS-->>S: map
    alt ui_context_stats_campaign_edit_page_enabled != true
        Note over S: fullStat.Previous = nil
    end

    S-->>H: *model.FullStat
    H-->>C: 200 JSON<br/>{advertId, begin, end, days, previous,<br/>nmStats, sideNmStats, appType}
```