FROM postgres:alpine

COPY scripts/init.sql /docker-entrypoint-initdb.d/

#COPY scripts/setup.sh ./

#RUN chmod a+r /docker-entrypoint-initdb.d/*

EXPOSE 5432

#CMD sh setup.sh