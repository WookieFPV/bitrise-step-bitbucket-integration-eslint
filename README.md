# bitbucket-integration-eslint

runs eslint and posts issues to bitbucket server


## How to use this Step

Can be run directly with the [bitrise CLI](https://github.com/bitrise-io/bitrise),
just `git clone` this repository, `cd` into it's folder in your Terminal/Command Line
and call `bitrise run test`.

*Check the `bitrise.yml` file for required inputs which have to be
added to your `.bitrise.secrets.yml` file!*

examples data based on url of your project: https://bitbucket.your-company.com/projects/FOO/repos/foo-app

```
- git::https://github.com/WookieFPV/bitrise-step-bitbucket-integration-eslint.git:
        inputs:
        - PROJECT_ID: "$PROJECT_ID"
        - REPORT_NAME: com.wookiefpv.eslintreporter
        - BITBUCKET_SERVER_URL: "$BITBUCKET_SERVER_URL"
        title: bitbucket ESLint integration
        is_always_run: true


Secrets: 
    BITBUCKET_SERVER_TOKEN = "examplePersonalToken" # (View Profile -> Manage Account -> HTTP access tokens with just READ permissions)

Env Vars:
    BITBUCKET_SERVER_URL = "https://bitbucket.your-company.com"
    PROJECT_ID = "FOO"



