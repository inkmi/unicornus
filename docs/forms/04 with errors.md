----
title: Displaying Errors
----

Unicornus can render form with errors. Errors are `map[string]string` and contain the field which created the error and an error text. These errors together with the data are rendered with `RenderFormWithErrors`.

![](cmd/example/example2.go, 1)

Results in

<img src="https://raw.githubusercontent.com/inkmi/unicornus/master/formexample.png" width="600">
