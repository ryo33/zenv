Feature: git
    Scenario: Add checkout
        Given the environment
        When I successfully run `zenv git checkout a`
        Then the file ".zenvl/git-checkout" should match /(.+) a/
    Scenario: Activate checkout
        Given the directory "foo"
        And I cd to "foo"
        And the environment
        And the git repository
        And the git branch "bar"
        And I run `zenv git checkout bar`
        And I cd to "."
        When I successfully run `git rev-parse --abbrev-ref HEAD`
        Then the output should contain exactly:
        """
        bar
        """
    Scenario: Deactivate checkout
        Given the directory "foo"
        And I cd to "foo"
        And the environment
        And the git repository
        And the git branch "bar"
        And I run `zenv git checkout bar`
        And I cd to "."
        And I cd to ".."
        When I successfully run `git --git-dir=foo/.git rev-parse --abbrev-ref HEAD`
        Then the output should contain exactly:
        """
        master
        """
