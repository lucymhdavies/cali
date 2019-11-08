# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

[Unreleased]: https://github.com/skybet/cali/compare/v0.4.0...master
## [Unreleased]


[0.4.0]:      https://github.com/skybet/cali/compare/v0.3.0...v0.4.0
## [0.4.0] - 2019-11-08

### Added
- Support for Go Modules - https://github.com/skybet/cali/pull/56
- Actively catching terminal resize - https://github.com/skybet/cali/pull/44

### Fixed
- Don't set WorkDir to a constant if it is already set - https://github.com/skybet/cali/pull/52


[0.3.0]:      https://github.com/skybet/cali/compare/v0.2.0...v0.3.0
## [0.3.0] - 2018-06-04

### Removed
- Remove auto-mount of aws credentials, from ~/.aws

### Changed
- expose cli.Name such that Cali apps can get their own name

[0.2.0]:      https://github.com/skybet/cali/compare/v0.1.2...v0.2.0
## [0.2.0] - 2018-04-11
### Changed
- Slight refactoring based on static analysis
- DockerClient.InitDocker does nothing if Docker client was previously initialised
- All DockerClient methods which rely on an initialised Docker client now run InitDocker. As a result, a few DockerClient methods now have the possibilities of returning errors: ContainerExists, ImageExists
- Auto-detects whether host process is a terminal, and uses that to automatically decide whether to run container non-interactive
- If running non-interactive, disable TTY on container

[0.1.2]:      https://github.com/skybet/cali/compare/v0.1.1...v0.1.2
## [0.1.2] - 2018-04-07
### Fixed
- PullImage now always pulls as should be expected
- StartContainer now only calls PullImage when the image does not exist locally
- Other miscelaneous refactoring

[0.1.1]:      https://github.com/skybet/cali/compare/v0.1.0...v0.1.1
## [0.1.1] - 2018-02-11
### Added
- This CHANGELOG file

### Changed
- Git data containers now have the repo name, branch and directory in the container name


[0.1.0]:      https://github.com/skybet/cali/compare/init...v0.1.0
## [0.1.0] - 2018-02-03
### Added
- Tagged a stable version of Cali which devs can pin to
