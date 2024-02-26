----
title: Nested Data
----

Data in Unicornus can be nested. A struct can have sub structs and those are rendered into HTML.
An embedded struct is best layouted with `AddGroup`. The name of the group is the name of the
embedded struct, in this case `Sub`. The label of the group is displayed as a header, the
description of the group is displayed for explanation.

The `AddGroup` is given a function `func(f *uni.FormLayout)`. Inside this function (kind of a callback)
the layout of the group is defined. The root is the sub struct of the group.


![](cmd/example/example3.go, 1)

