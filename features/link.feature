Feature: link command
    Scenario: Display links
        Given the environment
        And the file named ".zenvl/link" with:
        """
        foo=bar
        bar=baz
        baz=qux
        """
        When I successfully run `zenv link`
        Then the output should contain exactly:
        """
        foo bar
        bar baz
        baz qux
        """
    Scenario: Add new link
        Given the environment
        When I successfully run `zenv link foo bar`
        Then the file ".zenvl/link" should contain exactly:
        """
        foo=bar
        """
    Scenario: Remove links
        Given the environment
        And the file named ".zenvl/link" with:
        """
        foo=bar
        bar=baz
        baz=qux
        """
        When I successfully run `zenv link --remove foo bar`
        Then the file ".zenvl/link" should contain exactly:
        """
        baz=qux
        """
