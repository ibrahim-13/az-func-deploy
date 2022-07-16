# `az-func-deploy`
 - Deploy Azure Function Apps (C#) from Configuration
 - Deploy functions in a group with sets
 - Select which functions should be deployed
 - Config is updated on change, thus it will remember the choices next time

> Installation of **DotNet6.0** is required.
> Installation of **Azure CLI** or **Core Tools** is required based on configured method.

## Arguments
| Name    | Description                        |
| ------- | ---------------------------------- |
| `--cli` | Run directly as CLI instead of TUI |

## Configuration
File Name: **deploy.config.json**
Location:
 - Working Directory, or
 - Directory of the executing binary

> *If no configuration was found, the app will panic and create a sample configuration file in the working directory.*

## Configuration Sample
```json
{
    "Method": "azfunc",
    "Sets": [
        {
            "Name": "",
            "ResourceGroupName": "",
            "FuncInfos": [
                {
                    "FuncName": "",
                    "ProjectDir": "",
                    "ShouldRun": false
                }
            ]
        }
    ],
    "CurrentSet": 0
}
```

### JSON Fields
| Field             | Description                                 |
| ----------------- | ------------------------------------------- |
| Method            | Cli command for deployment (see next table) |
| Sets              | Functions in a Resource Group               |
| Name              | Name of the set (ex. DEV)                   |
| ResourceGroupName | Target Resource Group for the set           |
| FuncName          | Resource name of the Function App           |
| ProjectDir        | Location of the Project Directory           |
| ShouldRun         | If this Function should be deployed         |
| CurrentSet        | Index of the set that will be deployed 5    |

### Deployment Methods
| Method   | Description                                               | Documentation       |
| -------- | --------------------------------------------------------- | ------------------- |
| `azfunc` | Deploy with `az functionapp deployment source config-zip` | [Docs][azfunc-docs] |
| `azzip`  | Deploy with `az webap deploy`                             | [Docs][azzip-docs]  |
| `func`   | Deploy with `func azure functionapp publish`              | [Docs][func-docs]   | 

## Limitations
 - Can not create sets that with Functions from multiple Resource Group
 - Only works with **DotNet6.0**

## Libraries
 - [tview][tview-gh] - Terminal UI library

## Resources
 - [Azure Functions Overview][msdoc-az-func]
 - [tveiw][tview-gh]
 - [Deploy files to App Service][azzip-docs]
 - [az functionapp deployment source config-zip][azfunc-docs]
 - [Deploy project files][func-docs]

[msdoc-az-func]: https://docs.microsoft.com/en-us/azure/azure-functions/functions-overview
[tview-gh]: https://github.com/rivo/tview
[azzip-docs]: https://docs.microsoft.com/en-us/azure/app-service/deploy-zip?tabs=cli
[azfunc-docs]: https://docs.microsoft.com/en-us/cli/azure/functionapp/deployment/source?view=azure-cli-latest#az-functionapp-deployment-source-config-zip
[func-docs]: https://docs.microsoft.com/en-us/azure/azure-functions/functions-run-local?tabs=v4%2Cwindows%2Ccsharp%2Cportal%2Cbash#project-file-deployment