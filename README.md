# authn-jwt-gitlab

# DEPRECATED
[https://docs.cyberark.com/AAM-DAP/Latest/en/Content/Integrations/gitlab.htm](https://docs.cyberark.com/AAM-DAP/Latest/en/Content/Integrations/gitlab.htm)

## Description
This project creates a Docker image that includes a Go binary that can be used to authenticate a JWT token against Conjur Secrets Manager and retrieve a secret value.  Ubuntu, Alpine, and UBI-FIPS versions are available.  The secret value is returned to STDOUT and can be used in a GitLab CI pipeline.

## Badges
[![](https://img.shields.io/docker/pulls/nfmsjoeg/authn-jwt-gitlab)](https://hub.docker.com/r/nfmsjoeg/authn-jwt-gitlab) [![](https://img.shields.io/discord/802650809246154792)](https://discord.gg/J2Tcdg9tmk) [![](https://img.shields.io/reddit/subreddit-subscribers/cyberark?style=social)](https://reddit.com/r/cyberark) ![](https://img.shields.io/github/license/infamousjoeg/authn-jwt-gitlab)

## Requirements

* [Docker GitLab Runner](https://docs.gitlab.com/runner/install/docker.html)
* [Conjur Secrets Manager](https://www.conjur.org)
* Conjur Policies for authentication & authorization (authn & authz)
  * [authn-jwt Conjur Policy with GitLab Service ID](https://github.com/infamousjoeg/conjur-policies/tree/master/authn/authn-jwt-gitlab.yml)
  * [Conjur Policy to create identity for GitLab Repository](https://github.com/infamousjoeg/conjur-policies/blob/16f7375b604646a48b8b59ac9ddc011b6c8a08c6/ci/gitlab/root.yml#L45)
  * [Conjur Policy to grant GitLab Repository identity to use synchronized secrets from CyberArk Vault](https://github.com/infamousjoeg/conjur-policies/blob/84b451b5025fd1bb5fc86c601d172cb27da81b00/grants/grants_ci.yml#L41)
  * [Conjur Policy to grant GitLab Repository identity ability to authenticate using authn-jwt/gitlab web service](https://github.com/infamousjoeg/conjur-policies/blob/84b451b5025fd1bb5fc86c601d172cb27da81b00/grants/grants_authn.yml#L23)

## Usage

1. Choose your GitLab Runner Docker container image based on your desired OS.  The following images are available:
   * nfmsjoeg/authn-jwt-gitlab:ubuntu
   * nfmsjoeg/authn-jwt-gitlab:alpine
   * nfmsjoeg/authn-jwt-gitlab:ubi-fips
2. Once a GitLab Runner Docker container is decided upon, include it in your GitLab CI Pipeline file.  The following example is for the nfmsjoeg/authn-jwt-gitlab:ubuntu image:
```yaml
ubuntu:
    stage: test
    tags:
        - docker
    image: nfmsjoeg/authn-jwt-gitlab:ubuntu
```
3. Be sure to properly tag the job in the GitLab CI Pipeline file with the proper tag to run the job on the GitLab Runner Docker container.  This is done in the above example using the `tags` key.
4. Variables must be set in the GitLab CI Pipeline file for the GitLab Runner Docker container to consume.  Those environment variables are:
    * `CONJUR_APPLIANCE_URL`
    * `CONJUR_ACCOUNT`
    * `CONJUR_AUTHN_JWT_SERVICE_ID`
    * `CONJUR_AUTHN_JWT_TOKEN`
    * `CONJUR_SECRET_ID`
5. To use the binary in a job executing on the GitLab Runner Docker container, review the [example GitLab CI Pipeline script](.gitlab-ci.yml) in this repository.

### Example GitLab CI YAML File

```yaml
variables:
  CONJUR_APPLIANCE_URL: "https://conjur.joegarcia.dev"
  CONJUR_ACCOUNT: "cyberarkdemo"
  CONJUR_AUTHN_JWT_SERVICE_ID: "gitlab"
  CONJUR_AUTHN_JWT_TOKEN: "${CI_JOB_JWT}"

ubuntu:
  tags:
    - docker
  image: nfmsjoeg/authn-jwt-gitlab:ubuntu-dev
  script:
    - export TEST_USERNAME=$(CONJUR_SECRET_ID="SyncVault/LOB_CI/DemoSafe/DemoSafe-testuser4890/username" /authn-jwt-gitlab)
    - export TEST_PASSWORD=$(CONJUR_SECRET_ID="SyncVault/LOB_CI/DemoSafe/DemoSafe-testuser4890/password" /authn-jwt-gitlab)
    - env | grep TEST_

alpine:
  tags:
    - docker
  image: nfmsjoeg/authn-jwt-gitlab:alpine-dev
  script:
    - export TEST_USERNAME=$(CONJUR_SECRET_ID="SyncVault/LOB_CI/DemoSafe/DemoSafe-testuser4890/username" /authn-jwt-gitlab)
    - export TEST_PASSWORD=$(CONJUR_SECRET_ID="SyncVault/LOB_CI/DemoSafe/DemoSafe-testuser4890/password" /authn-jwt-gitlab)
    - env | grep TEST_

ubi-fips:
  stage: test
  tags:
    - docker
  image: nfmsjoeg/authn-jwt-gitlab:ubi-fips-dev
  script:
    - export TEST_USERNAME=$(CONJUR_SECRET_ID="SyncVault/LOB_CI/DemoSafe/DemoSafe-testuser4890/username" /authn-jwt-gitlab)
    - export TEST_PASSWORD=$(CONJUR_SECRET_ID="SyncVault/LOB_CI/DemoSafe/DemoSafe-testuser4890/password" /authn-jwt-gitlab)
    - env | grep TEST_
```

## Support
This is a community supported project.  For support, please file an issue in this repository.

## Contributing
If you would like to contribute to this project, please review the [CONTRIBUTING.md](CONTRIBUTING.md) file.

## License
This project is licensed under MIT - see the [LICENSE](LICENSE) file for details.
