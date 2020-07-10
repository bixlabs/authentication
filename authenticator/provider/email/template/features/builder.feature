Feature: build HTML templates for reset password
    In order to send emails with that templates
    As a builder of templates
    I need to build the default or a customice html templates

    Scenario: No template provided
        Given an empty enviroment variable 
        When the system sends an email
        Then the email should arrive with the default template

    Scenario: Correct template provided
        Given an correct environment variable
        When the system sends an email
        Then the email should arrive with the template provided
    
    Scenario: Incorrect template provided
        Given a wrong enviroment variable
        When the systems sends an email
        Then the email should arrive with the default template

