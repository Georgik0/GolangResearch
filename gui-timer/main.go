package main

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/speaker"
)

type TimerApp struct {
	minutesEntry   *widget.Entry
	secondsEntry   *widget.Entry
	startButton    *widget.Button
	stopButton     *widget.Button
	timerLabel     *widget.Label
	isRunning      bool
	ctx            context.Context
	cancel         context.CancelFunc
	totalSeconds   int
	currentSeconds int
}

// Генератор многотональных сигналов (как у грузовика)
type truckHorn struct {
	freqs  []float64
	amps   []float64
	sample float64
	sr     beep.SampleRate
}

func newTruckHorn(sr beep.SampleRate) *truckHorn {
	return &truckHorn{
		// Основные частоты для создания богатого звука грузовика
		freqs: []float64{220, 330, 440, 660, 880}, // Пентатоника в низком регистре
		amps:  []float64{0.6, 0.4, 0.8, 0.3, 0.2}, // Амплитуды для каждой частоты
		sr:    sr,
	}
}

func (h *truckHorn) Stream(samples [][2]float64) (n int, ok bool) {
	for i := range samples {
		var val float64 = 0
		
		// Смешиваем несколько частот для создания богатого звука
		for j, freq := range h.freqs {
			// Добавляем небольшую модуляцию для более живого звука
			modulation := 1 + 0.1*math.Sin(h.sample*2*math.Pi*5/float64(h.sr))
			val += h.amps[j] * math.Sin(h.sample*2*math.Pi*freq*modulation/float64(h.sr))
		}
		
		// Добавляем легкую перегрузку для более грубого звука
		if val > 0.7 {
			val = 0.7 + 0.3*math.Tanh((val-0.7)*3)
		} else if val < -0.7 {
			val = -0.7 - 0.3*math.Tanh((-val-0.7)*3)
		}
		
		samples[i][0] = val * 0.5 // Уменьшаем громкость
		samples[i][1] = val * 0.5
		h.sample++
	}
	return len(samples), true
}

func (h *truckHorn) Err() error {
	return nil
}

// Простой генератор синусоиды (оставляем для совместимости)
type sineWave struct {
	freq   float64
	sample float64
	sr     beep.SampleRate
}

func newSineWave(sr beep.SampleRate, freq float64) *sineWave {
	return &sineWave{
		freq: freq,
		sr:   sr,
	}
}

func (s *sineWave) Stream(samples [][2]float64) (n int, ok bool) {
	for i := range samples {
		val := math.Sin(s.sample * 2 * math.Pi * s.freq / float64(s.sr))
		samples[i][0] = val
		samples[i][1] = val
		s.sample++
	}
	return len(samples), true
}

func (s *sineWave) Err() error {
	return nil
}

func NewTimerApp() *TimerApp {
	app := &TimerApp{}
	app.setupUI()
	app.initAudio()
	return app
}

func (t *TimerApp) initAudio() {
	sr := beep.SampleRate(44100)
	speaker.Init(sr, sr.N(time.Second/10))
}

func (t *TimerApp) setupUI() {
	t.minutesEntry = widget.NewEntry()
	t.minutesEntry.SetPlaceHolder("Минуты")
	t.minutesEntry.SetText("1")

	t.secondsEntry = widget.NewEntry()
	t.secondsEntry.SetPlaceHolder("Секунды")
	t.secondsEntry.SetText("30")

	t.startButton = widget.NewButton("Старт", t.startTimer)
	t.stopButton = widget.NewButton("Стоп", t.stopTimer)
	t.stopButton.Disable()

	t.timerLabel = widget.NewLabel("00:00")
	t.timerLabel.Alignment = fyne.TextAlignCenter
	t.timerLabel.TextStyle = fyne.TextStyle{Bold: true}
}

func (t *TimerApp) startTimer() {
	if t.isRunning {
		return
	}

	minutesStr := t.minutesEntry.Text
	secondsStr := t.secondsEntry.Text

	minutes, err1 := strconv.Atoi(minutesStr)
	seconds, err2 := strconv.Atoi(secondsStr)

	if err1 != nil || err2 != nil || minutes < 0 || seconds < 0 {
		t.updateTimerLabel("Ошибка ввода!")
		return
	}

	t.totalSeconds = minutes*60 + seconds
	if t.totalSeconds == 0 {
		t.updateTimerLabel("Время должно быть > 0")
		return
	}

	t.isRunning = true
	t.ctx, t.cancel = context.WithCancel(context.Background())

	t.startButton.Disable()
	t.stopButton.Enable()
	t.minutesEntry.Disable()
	t.secondsEntry.Disable()

	go t.runTimer()
}

func (t *TimerApp) stopTimer() {
	if !t.isRunning {
		return
	}

	t.isRunning = false
	if t.cancel != nil {
		t.cancel()
	}

	t.startButton.Enable()
	t.stopButton.Disable()
	t.minutesEntry.Enable()
	t.secondsEntry.Enable()
	t.updateTimerLabel("Остановлено")
}

func (t *TimerApp) updateTimerLabel(text string) {
	fyne.Do(func() {
		t.timerLabel.SetText(text)
	})
}

func (t *TimerApp) runTimer() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for t.isRunning {
		t.currentSeconds = t.totalSeconds

		for t.currentSeconds > 0 && t.isRunning {
			select {
			case <-t.ctx.Done():
				return
			case <-ticker.C:
				minutes := t.currentSeconds / 60
				seconds := t.currentSeconds % 60
				timeStr := fmt.Sprintf("%02d:%02d", minutes, seconds)
				t.updateTimerLabel(timeStr)
				t.currentSeconds--
			}
		}

		if t.isRunning && t.currentSeconds <= 0 {
			t.playTruckHorn()
			t.updateTimerLabel("ВРЕМЯ!")
			select {
			case <-t.ctx.Done():
				return
			case <-time.After(1 * time.Second):
			}
		}
	}
}

func (t *TimerApp) playTruckHorn() {
	sr := beep.SampleRate(44100)
	
	// Создаем мощный звук грузовика
	horn := newTruckHorn(sr)
	
	// Делаем звук длиннее - 0.8 секунды
	hornSound := beep.Take(sr.N(time.Millisecond*800), horn)
	
	// Добавляем эффект реверберации для большей насыщенности
	reverb := effects.Echo(0.3, time.Millisecond*50)
	hornWithReverb := reverb(hornSound)
	
	// Воспроизводим звук
	speaker.Play(hornWithReverb)
}

func (t *TimerApp) getContainer() *fyne.Container {
	inputForm := container.NewGridWithColumns(2,
		widget.NewLabel("Минуты:"), t.minutesEntry,
		widget.NewLabel("Секунды:"), t.secondsEntry,
	)

	buttonContainer := container.NewGridWithColumns(2,
		t.startButton, t.stopButton,
	)

	mainContainer := container.NewVBox(
		widget.NewCard("Настройки таймера", "", inputForm),
		widget.NewSeparator(),
		widget.NewCard("", "", container.NewCenter(t.timerLabel)),
		widget.NewSeparator(),
		buttonContainer,
	)

	return mainContainer
}

func main() {
	myApp := app.New()
	myApp.SetIcon(nil)

	myWindow := myApp.NewWindow("Таймер")
	myWindow.Resize(fyne.NewSize(300, 280))
	myWindow.CenterOnScreen()

	timerApp := NewTimerApp()
	myWindow.SetContent(timerApp.getContainer())

	myWindow.SetCloseIntercept(func() {
		if timerApp.isRunning {
			timerApp.stopTimer()
		}
		myWindow.Close()
	})

	myWindow.ShowAndRun()
}

