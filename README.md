# go_azure_go

Some truly aweful code for brute forcing Azure resources unauthenticated.

This is unashamedly based on the Microburst (https://github.com/NetSPI/MicroBurst) EnumerateBlob and EnumerateSubdomain scripts, but those wrap powershell so are slow as all hell.
Tho is written in Go for naive concurrency in speed - but it will only be as good as the domains, keywords and wordlists that you put into it.

## Install
`sudo apt-get update && sudo apt-get -y install golang-go`

## Run
`go run brute.go`

## Usage
The script takes a list of common Azure endpoints (in `azure.txt`) and then uses a keyword (`base`, for example `evilcorp`) and a wordlist (`wordlist.txt`) to generate a list of possible Azure targets.
Then it tries to query those for DNS information.
Also - make sure you're querying a decent sized DNS server or your will get resolution failures (8.8.8.8 or 1.1.1.1 for example).
There's a manual delay in between requests to prevent overloading of DNS servers. You can adjust this downwards to improve performance.

## To Do
 - Make the script exit nicely rather than hanging
 - find a better way to throttle DNS queries rather than arbitrary delay numnbrs
 - increase rate by spreading requests between multiple DNS serveres
 - Add a starting request to retieve the target Tennet ID from a domain and then enumerate the associated domains (to harvets for keywords and SSO) the same way that `Invoke-AADIntReconAsOutsider -DomainName company.com | Format-Table` from  https://aadinternals.com/post/just-looking/ does
 - HTTP query any SSO/Auth endpoints to harvest tenant or subscription or aplication iD information from them.
 - Check for Oauth flows at
  - `https://login.windows.net/evilcorp.com/.well-known/openid-configuration`
  - `https://login.windows.net/<tenantid>/oauth2/token`
 - upgrade to include MFASweep endpoints:
  - `exchange.<domain.com>`
  - `mail.<domain.com>`
  - `webmail.<domain.com>`
  - `adfs.<domain.com>`
  - `lyncdiscover.domain.com`
  - `autodiscover.domain.com`
  - `<domain-com>.protection.outlook.com`
  - `<domain>.mail.onmicrosoft.com`
  - `<base>-prod-syd.azurewebsites.net`
  - Enumerate based on best pratice naming conventions dictated here: https://learn.microsoft.com/en-us/azure/cloud-adoption-framework/ready/azure-setup-guide/organize-resources

## Also Check the Following
 - Greyhatwarfare
 - Other better OSINT tools like BBOT
 - you might want to use the `permutations.txt` worklist from Microburst as a basic wordlist
 - looking into https://dnsdumpster.com/ or larger public DNS database (instead of just brute forcing) may also yield good info.

## FAQ
### Isn't this just a subdomain enumerator? Why did you use the better ones out there?
Mostly i wanted an excuse to learn some go an I have great difficulty quickly enumerating Azure resources for a tenant or subscription from an unauth perspective.
Use other better software for brute forcing subdomains. This just might help quickly discover public Azure resources that a client is using.

