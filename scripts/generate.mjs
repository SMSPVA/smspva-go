import { readFile, writeFile, mkdir } from 'node:fs/promises';
import { execFileSync } from 'node:child_process';
import { fileURLToPath } from 'node:url';
const root=new URL('../',import.meta.url);
const manifest=JSON.parse(await readFile(new URL('../contracts.json',import.meta.url),'utf8'));
await mkdir(root,{recursive:true});
await writeFile(new URL('contracts.json',root),JSON.stringify(manifest,null,2)+'\n');
await mkdir(new URL('testdata/',root),{recursive:true});
await writeFile(new URL('testdata/cases.json',root),await readFile(new URL('../testdata/cases.json',import.meta.url),'utf8'));
const cap=s=>s[0].toUpperCase()+s.slice(1);
const field=s=>({orderid:'OrderID',order_id:'OrderID',partnerkey:'PartnerKey',partner_key:'PartnerKey',id:'ID'}[s]??s.split(/[_-]/).map(cap).join(''));
let source='// Code generated from sdk/contracts.json. DO NOT EDIT.\npackage smspva\n\nimport "context"\n\n';
let tests='// Code generated from sdk/contracts.json. DO NOT EDIT.\npackage smspva\n\nimport ("context";"encoding/json")\n\nvar namedCalls=map[string]func(context.Context,*Client,json.RawMessage)(Response,error){\n';
let index='# Go SDK methods\n\nGenerated from the shared SDK contracts. All methods accept a context and optional request options. Request fields are pointers: nil omits a field. `Ptr(value)` creates a pointer.\n\n';
for(const family of ['activation','rental','email']) {
  const cls=cap(family);
  source+=`type ${cls} struct { client *Client }\n\n`;
  index+=`## ${cls}\n\n| Method | HTTP |\n| --- | --- |\n`;
  for(const [key,op] of Object.entries(manifest.operations).filter(([,v])=>v.family===family)) {
    const name=cap(op.jsName), type=cls+name+'Params';
    if(op.params.length) {
      source+=`type ${type} struct {\n`;
      for(const p of op.params) source+=`  ${field(p.name)} *${({string:'string',integer:'int64',number:'float64',boolean:'bool'})[p.schema.type]} ${JSON.stringify('json:"'+p.name+',omitempty"')} // ${p.required?'Required.':'Optional.'}\n`;
      source+='}\n\n';
    }
    source+=`// ${name} calls ${op.method} ${op.path}${op.action?'?method='+op.action:''}.\n`;
    if(op.stateful)source+='// This operation changes state or may charge the account. It is never retried by the SDK.\n';
    source+=`func (r *${cls}) ${name}(ctx context.Context${op.params.length?', params '+type:''}, options ...RequestOption) (Response,error) {\n return r.client.call(ctx,${JSON.stringify(key)},${op.params.length?'params':'nil'},options...)\n}\n\n`;
    tests+=`${JSON.stringify(key)}:func(ctx context.Context,c *Client,raw json.RawMessage)(Response,error){${op.params.length?`var p ${type};if err:=json.Unmarshal(raw,&p);err!=nil{return Response{},err};`:''}return c.${cls}.${name}(ctx${op.params.length?',p':''})},\n`;
    index+=`| \`client.${cls}.${name}\` | \`${op.method} ${op.path}${op.action?'?method='+op.action:''}\` |\n`;
  }
  index+='\n';
}
tests+='}\n';
await writeFile(new URL('resources.go',root),source);
await writeFile(new URL('named_calls_test.go',root),tests);
await writeFile(new URL('METHODS.md',root),index);
execFileSync(process.env.GOFMT ?? 'gofmt',['-w',fileURLToPath(new URL('resources.go',root)),fileURLToPath(new URL('named_calls_test.go',root))],{stdio:'inherit',windowsHide:true});
console.log('Generated Go resources and 56 typed method test adapters.');
