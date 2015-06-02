# Skeleton Service Cookbook

> Cookbook for deploying the skeleton-service

## Requirements

* [ChefDK](https://downloads.chef.io/chef-dk/)
* [Vagrant](https://www.vagrantup.com/)
* [VirtualBox](https://www.virtualbox.org/wiki/Downloads)

## Installation

1. Install requirements listed above
2. ```./go idea``` to setup the Intellij project
3. ```./go precommit``` to validate the project is setup correctly

## Usage

    ./go install        Install dependant cookbooks via Berkshelf
    ./go lint           Run linting on cookbooks via Foodcritic, with default ruleset
    ./go converge       Spin up local virtual machine and apply cookbooks
    ./go verify         Run integration tests against the local virtual machine
    ./go destroy        Destroy the local virtual machine
    ./go precommit      Run all validations before pushing code

## Acceptance Tests

This cookbook has a suite of acceptance tests covering the main functionality of the cookbook.

If you modify the cookbook and want to verify your changes (and that you haven't broken anything else), run the tests.

To run all validations before pushing code, run ```./go precommit```

## Author

Jen Scobie (jenscobie@gmail.com)