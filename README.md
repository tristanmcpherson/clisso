# Clisso: CLI Single Sign-On

**WIP Warning! This project is still under development and isn't expected
to work yet.**

*Clisso* allows you to retrieve temporary credentials for cloud
providers and APIs by authenticating with an identity provider (IdP).

The following cloud providers are currently supported:

- [AWS](https://aws.amazon.com/)

The following identity providers are currently supported:

- [OneLogin](https://www.onelogin.com/)

## Installation

1. Inside `$GOPATH/src/github.com/johananl/clisso` run `dep ensure` to install dependencies.
1. Run `go install`. This will put the `clisso` binary in your `$GOPATH/bin` directory.

## Configuration

Create a file called `.clisso.yaml` in your home directory. Following is a
sample configuration:

    providers:
      onelogin:
        clientSecret: xxxxxxxx
        clientId: xxxxxxxx
        subdomain: mydomain
    apps:
      dev-account:
        provider: onelogin
        appId: 123456
        principalArn: arn:aws:iam::000000000000:saml-provider/My-SAML-IdP
        roleArn: arn:aws:iam::000000000000:role/My-IAM-Role
      prod-account:
        provider: onelogin
        appId: 234567
        principalArn: arn:aws:iam::000000000000:saml-provider/My-SAML-IdP
        roleArn: arn:aws:iam::000000000000:role/My-IAM-Role


## Usage

Run `clisso get <app-name>` and enter your username, password and OTP
to get temporary credentials.
