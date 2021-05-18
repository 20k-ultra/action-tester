# Action Tester
zzz
This repo uses 2 actions to automate balenaCloud releases:

 - [Build Release](.github/workflows/pull_request.yml)
 - [Deploy Release](.github/workflows/master_merge.yml)

 ### Build Release action
 ---
 When: This action is configured to run on all pull requests.

 Why: Builds a release from source marked as a draft.

### Deploy Release action
---
 When: This action will run once a pull request is closed and merges into master.

 Why: Finds a previously built release and marks it as final.