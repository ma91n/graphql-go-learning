
```bash
# 初回だけ
docker volume create pg-data-eki

# 起動
docker-compose up --build

# 停止
docker-compose down
```



```bash
# 駅名検索
xo pgsql://postgres:postgres@localhost/postgres?sslmode=disable -N -M -B -T StationByName -o models/ << ENDSQL
select l.line_cd, l.line_name, s.station_cd, station_g_cd, s.station_name, s.address 
from station s
         inner join line l on s.line_cd = l.line_cd
where s.station_name = %%stationName string%%
  and s.e_status = 0
ENDSQL

# 駅CD検索
xo pgsql://postgres:postgres@localhost/postgres?sslmode=disable -N -M -B -T StationByCD -o models/ << ENDSQL
select l.line_cd, l.line_name, s.station_cd, station_g_cd, s.station_name, s.address 
from station s
         inner join line l on s.line_cd = l.line_cd
where s.station_cd = %%stationCD int%%
  and s.e_status = 0
ENDSQL

# 乗換駅検索
xo pgsql://postgres:postgres@localhost/postgres?sslmode=disable -N -M -B -T Transfer -o models/ << ENDSQL
select s.station_cd,
       ls.line_cd,
       ls.line_name,
       s.station_name,
       s.station_g_cd,
       s.address,
       COALESCE(lt.line_cd, 0)     as transfer_line_cd,
       COALESCE(lt.line_name, '')   as transfer_line_name,
       COALESCE(t.station_cd, 0)   as transfer_station_cd,
       COALESCE(t.station_name, '') as transfer_station_name,
       COALESCE(t.address, '')      as transfer_address
from station s
         left outer join station t on s.station_g_cd = t.station_g_cd and s.station_cd <> t.station_cd
         left outer join line ls on s.line_cd = ls.line_cd
         left outer join line lt on t.line_cd = lt.line_cd
where s.station_cd = %%stationCD int%%
ENDSQL

xo pgsql://postgres:postgres@localhost/postgres?sslmode=disable -N -M -B -T Before -o models/ << ENDSQL
select sl.line_cd,
        sl.line_name,
        s.station_cd,
        s.station_name,
        s.address,
        COALESCE(js.station_cd, 0)    as before_station_cd,
        COALESCE(js.station_name, '') as before_station_name,
        COALESCE(js.station_g_cd, 0)  as before_station_g_cd,
        COALESCE(js.address, '')      as before_station_address
 from station s
          left outer join line sl on s.line_cd = sl.line_cd
          left outer join station_join j on s.line_cd = j.line_cd and s.station_cd = j.station_cd1
          left outer join station js on j.station_cd2 = js.station_cd
 where s.e_status = 0
   and s.station_cd = %%stationCD int%%
ENDSQL

xo pgsql://postgres:postgres@localhost/postgres?sslmode=disable -N -M -B -T After -o models/ << ENDSQL
select sl.line_cd,
        sl.line_name,
        s.station_cd,
        s.station_name,
        s.address,
        COALESCE(js.station_cd, 0)    as after_station_cd,
        COALESCE(js.station_name, '') as after_station_name,
        COALESCE(js.station_g_cd, 0)  as after_station_g_cd,
        COALESCE(js.address, '')      as after_station_address
 from station s
          left outer join line sl on s.line_cd = sl.line_cd
          left outer join station_join j on s.line_cd = j.line_cd and s.station_cd = j.station_cd2
          left outer join station js on j.station_cd1 = js.station_cd
 where s.e_status = 0
   and s.station_cd = %%stationCD int%%
ENDSQL
```