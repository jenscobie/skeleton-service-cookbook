# Skeleton Service Cookbook

> Cookbook for deploying the skeleton-service

## Requirements

* VirtualBox
* Chef Development Kit

## Installation

1. Install requirements listed above
2. ```./go idea``` to setup the Intellij project
3. ```./go precommit``` to validate the project is setup correctly

## Usage

    ./go lint           Run foodcritic linting on cookbooks, with default rules
    ./go converge       Spin up local VM and apply cookbooks
    ./go verify         Run integration tests against local VM
    ./go destroy        Destroy the local VM
    ./go precommit      Run all validations before pushing code

## Acceptance Tests

This cookbook has a suite of acceptance tests covering the main functionality of the cookbook.

If you modify the cookbook and want to verify your changes (and that you haven't broken anything else), run the tests.

To run all validations before pushing code, run ```./go precommit```

## Author

Jen Scobie (jenscobie@gmail.com)