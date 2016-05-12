This repository is used to demonstrate to Google Engineering that their Health
Checker installer is broken in their default Flexible VM environment.

To deploy, you must use aedeploy thusly:

        aedeploy gcloud --quiet --verbosity debug preview app deploy --no-promote app.yaml
