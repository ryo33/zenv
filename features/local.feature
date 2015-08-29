Feature: List of activated environment
    Scenario: Create local environment
        When I successfully run `zenv local`
        Then the file ".zenvl/info" should contain:
        """
        name=/Users/ryo/dev/go/src/github.com/ryo33/zenv/tmp/aruba
        dir=/Users/ryo/dev/go/src/github.com/ryo33/zenv/tmp/aruba/.zenvl
        global=false
        recursive=true
        exclusive=false
        """
    Scenario: Create not-recursive local environment
        When I successfully run `zenv local --not-recursive`
        Then the file ".zenvl/info" should contain:
        """
        name=/Users/ryo/dev/go/src/github.com/ryo33/zenv/tmp/aruba
        dir=/Users/ryo/dev/go/src/github.com/ryo33/zenv/tmp/aruba/.zenvl
        global=false
        recursive=false
        exclusive=false
        """
    Scenario: Create exclusive local environment
        When I successfully run `zenv local --exclusive`
        Then the file ".zenvl/info" should contain:
        """
        name=/Users/ryo/dev/go/src/github.com/ryo33/zenv/tmp/aruba
        dir=/Users/ryo/dev/go/src/github.com/ryo33/zenv/tmp/aruba/.zenvl
        global=false
        recursive=true
        exclusive=true
        """
