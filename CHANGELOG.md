# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [v0.17.0] - 2024-09-02
### :sparkles: New Features
- [`90551e8`](https://github.com/mole-squad/soq-api/commit/90551e86043c514b24641d7706ab9b8ebb6a81b9) - validate name on focus area create *(PR [#47](https://github.com/mole-squad/soq-api/pull/47) by [@burkel24](https://github.com/burkel24))*


## [v0.16.0] - 2024-07-21
### :sparkles: New Features
- [`5558526`](https://github.com/mole-squad/soq-api/commit/5558526ef82b401597d817d9240caee8180bad96) - return 201 for login *(PR [#46](https://github.com/mole-squad/soq-api/pull/46) by [@burkel24](https://github.com/burkel24))*


## [v0.15.0] - 2024-07-21
### :sparkles: New Features
- [`7befb30`](https://github.com/mole-squad/soq-api/commit/7befb3027fc4e578bf7e73c139a7771cd0c1ae9c) - use patch instead of post for resolve task *(PR [#45](https://github.com/mole-squad/soq-api/pull/45) by [@burkel24](https://github.com/burkel24))*


## [v0.14.0] - 2024-07-18
### :sparkles: New Features
- [`f47a63f`](https://github.com/mole-squad/soq-api/commit/f47a63fdc95a19461def398ee450340253902656) - use generic resource for quotas *(PR [#42](https://github.com/mole-squad/soq-api/pull/42) by [@burkel24](https://github.com/burkel24))*
- [`199eade`](https://github.com/mole-squad/soq-api/commit/199eade2373ca8699c5893309da961971e749962) - port devices to generic resource *(PR [#43](https://github.com/mole-squad/soq-api/pull/43) by [@burkel24](https://github.com/burkel24))*
- [`21516d9`](https://github.com/mole-squad/soq-api/commit/21516d9031c6b5a6fe232f77ee95445cedc273aa) - add timewindows APIs *(PR [#44](https://github.com/mole-squad/soq-api/pull/44) by [@burkel24](https://github.com/burkel24))*


## [v0.13.0] - 2024-07-18
### :sparkles: New Features
- [`b0c272c`](https://github.com/mole-squad/soq-api/commit/b0c272c1573b75e168adbff3dd572e0477c94609) - port focusareas to generic resource *(PR [#41](https://github.com/mole-squad/soq-api/pull/41) by [@burkel24](https://github.com/burkel24))*


## [v0.12.0] - 2024-07-17
### :sparkles: New Features
- [`4814b38`](https://github.com/mole-squad/soq-api/commit/4814b38213dd1732e97d69c3640a81e1c22da1a4) - consolidate generic patterns *(PR [#40](https://github.com/mole-squad/soq-api/pull/40) by [@burkel24](https://github.com/burkel24))*


## [v0.11.0] - 2024-07-17
### :sparkles: New Features
- [`dcbea23`](https://github.com/mole-squad/soq-api/commit/dcbea232c56b52e964edd2e022588baccb6a28d5) - create generic resource service *(PR [#39](https://github.com/mole-squad/soq-api/pull/39) by [@burkel24](https://github.com/burkel24))*


## [v0.10.0] - 2024-07-17
### :sparkles: New Features
- [`0780e52`](https://github.com/mole-squad/soq-api/commit/0780e521b17b98b4b3e6c6259130afe48c9118d5) - move GetUserFromCtx to AuthService *(PR [#38](https://github.com/mole-squad/soq-api/pull/38) by [@burkel24](https://github.com/burkel24))*


## [v0.9.0] - 2024-07-17
### :sparkles: New Features
- [`fb7f53f`](https://github.com/mole-squad/soq-api/commit/fb7f53ff5b5f9dec56be11098502ab6f7df2c234) - create generic repo *(PR [#37](https://github.com/mole-squad/soq-api/pull/37) by [@burkel24](https://github.com/burkel24))*


## [v0.8.2] - 2024-07-17
### :bug: Bug Fixes
- [`d868880`](https://github.com/mole-squad/soq-api/commit/d868880f2427936909118c336c9d2578b052c1f5) - dont exit on agenda send error *(PR [#36](https://github.com/mole-squad/soq-api/pull/36) by [@burkel24](https://github.com/burkel24))*


## [v0.8.1] - 2024-07-17
### :bug: Bug Fixes
- [`2d32141`](https://github.com/mole-squad/soq-api/commit/2d321412ed18fb4c3a3edf1c23d51285d7fa119e) - handle empty agenda in send cmd *(PR [#35](https://github.com/mole-squad/soq-api/pull/35) by [@burkel24](https://github.com/burkel24))*


## [v0.8.0] - 2024-07-17
### :sparkles: New Features
- [`9904e94`](https://github.com/mole-squad/soq-api/commit/9904e947bca7b1547094df79c094698677f35330) - add generic rest controller *(PR [#34](https://github.com/mole-squad/soq-api/pull/34) by [@burkel24](https://github.com/burkel24))*


## [v0.7.0] - 2024-07-16
### :sparkles: New Features
- [`cde645f`](https://github.com/mole-squad/soq-api/commit/cde645fdd0ccffea4e57009a187cdf5f8aa3a164) - add device CRUD apis *(PR [#33](https://github.com/mole-squad/soq-api/pull/33) by [@burkel24](https://github.com/burkel24))*

### :wrench: Chores
- [`70dfefb`](https://github.com/mole-squad/soq-api/commit/70dfefbdecb1bbc426d48d123872d25115ecacdd) - refactor task controller to use context middleware *(PR [#32](https://github.com/mole-squad/soq-api/pull/32) by [@burkel24](https://github.com/burkel24))*


## [v0.6.0] - 2024-07-15
### :sparkles: New Features
- [`0d85c81`](https://github.com/mole-squad/soq-api/commit/0d85c81e3158996065942f709862b0cde9fafed1) - add resolve task api *(PR [#31](https://github.com/mole-squad/soq-api/pull/31) by [@burkel24](https://github.com/burkel24))*


## [v0.5.0] - 2024-07-13
### :sparkles: New Features
- [`7d244f8`](https://github.com/mole-squad/soq-api/commit/7d244f839b1a58d3b80dec71525add6420f7af5b) - add create user API *(PR [#30](https://github.com/mole-squad/soq-api/pull/30) by [@burkel24](https://github.com/burkel24))*


## [v0.4.0] - 2024-07-13
### :sparkles: New Features
- [`75b5e03`](https://github.com/mole-squad/soq-api/commit/75b5e035fd855da46b72c79561c4b1ed14292816) - set ref for release checkout *(commit by [@burkel24](https://github.com/burkel24))*


## [v0.2.8] - 2024-07-13
### :bug: Bug Fixes
- [`9b5489b`](https://github.com/mole-squad/soq-api/commit/9b5489b2df8427d46824238d31d4c6b775480388) - set config on web dyno *(commit by [@burkel24](https://github.com/burkel24))*


## [v0.2.7] - 2024-07-13
### :bug: Bug Fixes
- [`4794c30`](https://github.com/mole-squad/soq-api/commit/4794c300685ca19cb6632085d7fdad7d97d1c3cb) - add tzinfo to docker image + cleanup *(PR [#29](https://github.com/mole-squad/soq-api/pull/29) by [@burkel24](https://github.com/burkel24))*

### :wrench: Chores
- [`979eb4e`](https://github.com/mole-squad/soq-api/commit/979eb4e246df21b08873e4d98f0d4da210349ae7) - rename ci workflow [skip ci] *(commit by [@burkel24](https://github.com/burkel24))*


## [v0.2.6] - 2024-07-13
### :bug: Bug Fixes
- [`6a1fd55`](https://github.com/mole-squad/soq-api/commit/6a1fd55bfdf032e70169d0d70b103e2085668cc3) - tag release *(commit by [@burkel24](https://github.com/burkel24))*

[v0.2.6]: https://github.com/mole-squad/soq-api/compare/v0.2.5...v0.2.6
[v0.2.7]: https://github.com/mole-squad/soq-api/compare/v0.2.6...v0.2.7
[v0.2.8]: https://github.com/mole-squad/soq-api/compare/v0.2.7...v0.2.8
[v0.4.0]: https://github.com/mole-squad/soq-api/compare/v0.3.0...v0.4.0
[v0.5.0]: https://github.com/mole-squad/soq-api/compare/v0.4.0...v0.5.0
[v0.6.0]: https://github.com/mole-squad/soq-api/compare/v0.5.0...v0.6.0
[v0.7.0]: https://github.com/mole-squad/soq-api/compare/v0.6.0...v0.7.0
[v0.8.0]: https://github.com/mole-squad/soq-api/compare/v0.7.0...v0.8.0
[v0.8.1]: https://github.com/mole-squad/soq-api/compare/v0.8.0...v0.8.1
[v0.8.2]: https://github.com/mole-squad/soq-api/compare/v0.8.1...v0.8.2
[v0.9.0]: https://github.com/mole-squad/soq-api/compare/v0.8.2...v0.9.0
[v0.10.0]: https://github.com/mole-squad/soq-api/compare/v0.9.0...v0.10.0
[v0.11.0]: https://github.com/mole-squad/soq-api/compare/v0.10.0...v0.11.0
[v0.12.0]: https://github.com/mole-squad/soq-api/compare/v0.11.0...v0.12.0
[v0.13.0]: https://github.com/mole-squad/soq-api/compare/v0.12.0...v0.13.0
[v0.14.0]: https://github.com/mole-squad/soq-api/compare/v0.13.0...v0.14.0
[v0.15.0]: https://github.com/mole-squad/soq-api/compare/v0.14.0...v0.15.0
[v0.16.0]: https://github.com/mole-squad/soq-api/compare/v0.15.0...v0.16.0
[v0.17.0]: https://github.com/mole-squad/soq-api/compare/v0.16.0...v0.17.0
