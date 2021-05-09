# 711-location

This app helps to find out geo-location to the 711 petrol stations in Australia
with the best fuel prices. It has been a pain for me to search GoogleMap for GPS
coordinates, which gets consumed in GPS location mocking apps for *locking*
desired fuel price.

## How to use

- Default case `./711-location`
- To override default fuel type which is *U91*: set environment variable `FUEL_TYPE`.
i.e. `FUEL_TYPE=U95 ./711-location`
- Available fuel types to choose from:
    - E10
    - U91
    - U95
    - U98
    - Diesel
    - LPG


## How to compile your own executables

This repo has included executables ready for use:
- 711-location for Mac
- 711-location.exe for Windows

However, it is fine if you
want to take precautions and only run executables that you built for yourself.

i.e. for Windows `GOOS=windows go build`
i.e. for Mac     `GOOS=darwin go build`