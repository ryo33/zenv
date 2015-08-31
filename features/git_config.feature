Feature: git-config
    Scenario: Add config
        Given the environment
        When I successfully run `zenv git config a b`
        Then the file ".zenvl/git-config" should match /(.+) a b/
    Scenario: Add global config
        Given the environment
        When I successfully run `zenv git config a b`
        Then the file ".zenvl/git-config" should match /a b/
    Scenario: Remove config
        Given the environment
        And I successfully run `zenv git config foo bar`
        When I successfully run `zenv git config --remove foo`
        Then the file ".zenvl/git-config" should not exist
    Scenario: Activate config
        Given the directory "foo"
        And I cd to "foo"
        And the environment
        And the git repository
        And the git branch "bar"
        And I successfully run `zenv git config user.name bar`
        And I cd to "."
        When I successfully run `git config user.name`
        Then the output should contain exactly:
        """
        bar
        """
    Scenario: Deactivate config
        Given the directory "foo"
        And I cd to "foo"
        And the environment
        And the git repository
        And the git branch "bar"
        And I successfully run `zenv git config user.name bar`
        And I cd to "."
        And I cd to ".."
        When I successfully run `git --git-dir=./foo/.git config user.name`
        Then the output from "git --git-dir=./foo/.git config user.name" should contain exactly:
        """
        """
    Scenario: Activate config another git directory
        Given the directory "foo"
        And I cd to "foo"
        And the git repository
        And the git branch "bar"
        And I cd to ".."
        And the environment
        And I successfully run `zenv git config --directory ./foo user.name bar`
        And I cd to "."
        When I successfully run `git --git-dir=./foo/.git config user.name`
        Then the output should contain exactly:
        """
        bar
        """
    Scenario: Deactivate config another git directory
        Given the directory "foo"
        And I cd to "foo"
        And the git repository
        And the git branch "bar"
        And I cd to ".."
        And the not-recursive environment
        And I successfully run `zenv git config --directory ./foo user.name bar`
        And I cd to "."
        And I cd to "foo"
        When I successfully run `git config user.name`
        Then the output from "git config user.name" should contain exactly:
        """
        """
