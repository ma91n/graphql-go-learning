
```bash
# 初回だけ
docker volume create pg-data-eki

# 起動
docker-compose up --build

# 停止
docker-compose down
```



```bash
xo pgsql://postgres:postgres@localhost/postgres?sslmode=disable -N -M -B -T Station -o models/ << ENDSQL
select l.line_cd, l.line_name, s.station_cd, station_g_cd, s.station_name, s.address 
from station s
         inner join line l on s.line_cd = l.line_cd
where s.station_name = %%stationName string%%
  and s.e_status = 0
ENDSQL

xo pgsql://postgres:postgres@localhost/postgres?sslmode=disable -N -M -B -T StationConn -o models/ << ENDSQL
select li   .line_name,
       li.line_name_h,
       li.line_cd,
       st.station_cd,
       st.station_g_cd,
       st.address,
       st.station_name,
       COALESCE(st2.station_cd, 0)   as before_station_cd,
       COALESCE(st2.station_name, '') as before_station_name,
       COALESCE(st3.station_cd, 0)   as after_station_cd,
       COALESCE(st3.station_name, '') as after_station_name,
       COALESCE(gli.line_name, '')    as transfer_line_name,
       COALESCE(gs.station_cd, 0)    as transfer_station_cd,
       COALESCE(gs.station_name, '')  as transfer_station_name
from station st
         inner join line li on st.line_cd = li.line_cd
         left outer join station_join sjb on st.line_cd = sjb.line_cd and st.station_cd = sjb.station_cd2 
         left outer join station_join sja on st.line_cd = sja.line_cd and st.station_cd = sja.station_cd1
         left outer join station st2 on sjb.line_cd = st2.line_cd and sjb.station_cd1 = st2.station_cd
         left outer join station st3 on sja.line_cd = st3.line_cd and sja.station_cd2 = st3.station_cd
         left outer join station gs on st.station_g_cd = gs.station_g_cd and st.station_cd <> gs.station_cd
         left outer join line gli on gs.line_cd = gli.line_cd
where st.station_cd = %%stationCD int%%
  and st.e_status = 0
order by st.e_sort
ENDSQL
```