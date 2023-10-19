# bodycomposition-webapi

Bodycomposition-WebAPI is RESTful api to manage body composition data in
Garmin Connect Cloud (https://connect.garmin.com) from Web API

# OUTDATED - See a new version

New working alternative (19.10.2023) (https://github.com/lswiderski/yet-another-garmin-connect-client)

## Inspiration

Inspired by [bodycomposition cli version](https://github.com/davidkroell/bodycomposition).

## Use case

- Run it in docker or Serverless function as middleware for importing values from your scale to Garmin Connect Cloud.

- Used in Mobile App: https://github.com/lswiderski/mi-scale-exporter

## Installation

- DockerHub

  ```dockerfile
  docker pull lswiderski/bodycomposition-webapi:latest
  ```

- Source code build
  ```dockerfile
  docker build .
  ```

## Usage

- See [api.rest](https://github.com/lswiderski/bodycomposition-webapi/blob/main/api.rest) file

- Endpoint
  ```http
  /upload
  ```
- payload:

  ```json
  {
    "timeStamp": -1, // -1 for Time Now, or UNIX timestamp
    "weight": 67, // kg
    "percentFat": 14.9, // %
    "percentHydration": 58.4, // %
    "boneMass": 3.3, // kg
    "muscleMass": 55.1, // kg
    "visceralFatRating": 7,
    "physiqueRating": 5,
    "metabolicAge": 25,
    "bodyMassIndex": 20.8,
    "email": "{{email}}",
    "password": "{{password}}"
  }
  ```

## Coffee

<a href="https://www.buymeacoffee.com/lukaszswiderski" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/default-orange.png" alt="Buy Me A Coffee" height="41" width="174"></a>
