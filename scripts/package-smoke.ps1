param([string]$GoExecutable = 'go')
$ErrorActionPreference = 'Stop'
$moduleRoot = [IO.Path]::GetFullPath((Join-Path $PSScriptRoot '..'))
$checkRoot = Join-Path $moduleRoot '.checks'
New-Item -ItemType Directory -Force $checkRoot | Out-Null
$archive = Join-Path $checkRoot 'smspva-go-0.1.0-alpha.1.zip'
$sources = @('LICENSE','doc.go','go.mod','contracts.json','README.md','METHODS.md','client.go','errors.go','polling.go','resources.go','client_test.go','named_calls_test.go','testdata') | ForEach-Object { Join-Path $moduleRoot $_ }
Compress-Archive -LiteralPath $sources -DestinationPath $archive -Force
$smokeRoot = Join-Path $checkRoot ('install-' + [guid]::NewGuid().ToString('N'))
$installed = Join-Path $smokeRoot 'module'
Expand-Archive -LiteralPath $archive -DestinationPath $installed
& $GoExecutable -C $installed test ./...
if ($LASTEXITCODE -ne 0) { throw 'Tests from the standalone module archive failed' }
$consumer = Join-Path $smokeRoot 'consumer'
New-Item -ItemType Directory -Force $consumer | Out-Null
$modulePath = $installed.Replace('\','/')
$goMod = "module example.com/smspva-consumer`n`ngo 1.22`n`nrequire github.com/SMSPVA/smspva-go v0.0.0`nreplace github.com/SMSPVA/smspva-go => `"$modulePath`"`n"
[IO.File]::WriteAllText((Join-Path $consumer 'go.mod'), $goMod)
$consumerSource = @'
package main
import (
 "context"
 "net/http"
 "net/http/httptest"
 "strings"
 smspva "github.com/SMSPVA/smspva-go"
)
func main() {
 server:=httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
  if strings.HasPrefix(r.URL.Path,"/activation"){w.Write([]byte(`{"statusCode":200,"data":{"balance":"10.50"}}`))}else{w.Write([]byte(`{"status":1,"data":[]}`))}
 }));defer server.Close()
 c,err:=smspva.New(smspva.Options{APIKey:"local-smoke-key",ActivationBaseURL:server.URL,RentalBaseURL:server.URL,EmailBaseURL:server.URL});if err!=nil{panic(err)};defer c.CloseIdleConnections()
 ctx:=context.Background()
 if _,err=c.Activation.GetBalance(ctx);err!=nil{panic(err)}
 if _,err=c.Rental.GetCountries(ctx);err!=nil{panic(err)}
 if _,err=c.Email.GetProviders(ctx);err!=nil{panic(err)}
}
'@
[IO.File]::WriteAllText((Join-Path $consumer 'main.go'), $consumerSource)
& $GoExecutable -C $consumer run .
if ($LASTEXITCODE -ne 0) { throw 'Standalone Go consumer failed' }
Write-Output "Standalone archive and three-family consumer passed: $archive"
