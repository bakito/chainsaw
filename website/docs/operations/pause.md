# Pause

The `Pause` operation provides a means to add a pause to the executions.

## Configuration

The full structure of `Pause` is documented [here](../reference/apis/chainsaw.v1alpha1.md#chainsaw-kyverno-io-v1alpha1-Sleep).

### Features

| Supported features                                 |                    |
|----------------------------------------------------|:------------------:|
| [Bindings](../general/bindings.md) support         | :x:                |
| [Outputs](../general/outputs.md) support           | :x:                |
| [Templating](../general/templating.md) support     | :x:                |
| [Operation checks](../general/checks.md) support   | :x:                |

## Examples

```yaml
apiVersion: chainsaw.kyverno.io/v1alpha1
kind: Test
metadata:
  name: example
spec:
  steps:
  - try:
    - pause: {}
```
