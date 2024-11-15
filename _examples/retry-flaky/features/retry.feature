Feature: godog should be able to retry flaky tests
        Developers can define an amount of retries for godog to automatically retry it's scenario when they failed
        When a step in a scenario failed, godog will retry that scenario, including the scenario hooks and step hooks

    Scenario: retry flaky test at maximum 3 times before return error
        Given a step randomly failed
        When godog retry the scenario
        Then the scenario must pass