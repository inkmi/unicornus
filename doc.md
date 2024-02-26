
## Idea

The idea of Unicornus is to combine a data model in Go described as structs with validation tags
with a description of the form layout in Go to render an HTML form.


```
           ┌─────────────┐             ┌─────────────┐
           │             │             │             │
Validation │ Data Model  ├──────┬──────┤ Form Layout │
           │             │      │      │             │
           │             │      │      │             │
           └─────────────┘      │      └─────────────┘
                  ▲             │
                  │             │
                  │             │
                  │             │
                  │             │
                  │             │
                  │             ▼
                  │       ┌───────────┐
                  │       │           │
           Submit │       │           │
                  │       │   HTML    │
                  └───────┤   Form    │
                          │           │
                          │           │
                          └───────────┘
```


## Code Example

<pre tabindex="0" style="color:#e5e5e5;background-color:#000;"><code><span style="display:flex;"><span><span style="color:#fff;font-weight:bold">import</span> (
</span></span><span style="display:flex;"><span>	<span style="color:#0ff;font-weight:bold">&#34;fmt&#34;</span>
</span></span><span style="display:flex;"><span>	uni <span style="color:#0ff;font-weight:bold">&#34;github.com/inkmi/unicornus/pkg&#34;</span>
</span></span><span style="display:flex;"><span>	<span style="color:#0ff;font-weight:bold">&#34;net/http&#34;</span>
</span></span><span style="display:flex;"><span>)
</span></span><span style="display:flex;"><span>
</span></span><span style="display:flex;"><span><span style="color:#fff;font-weight:bold">type</span> errorexample <span style="color:#fff;font-weight:bold">struct</span> {
</span></span><span style="display:flex;"><span>	Name <span style="color:#fff;font-weight:bold">string</span>
</span></span><span style="display:flex;"><span>}
</span></span><span style="display:flex;"><span>
</span></span><span style="display:flex;"><span>	<span style="color:#007f7f">// The data of the form
</span></span></span><span style="display:flex;"><span><span style="color:#007f7f"></span>	d := errorexample{
</span></span><span style="display:flex;"><span>		Name: <span style="color:#0ff;font-weight:bold">&#34;Unicornus&#34;</span>,
</span></span><span style="display:flex;"><span>	}
</span></span><span style="display:flex;"><span>	<span style="color:#007f7f">// Create a FormLayout
</span></span></span><span style="display:flex;"><span><span style="color:#007f7f"></span>	<span style="color:#007f7f">// describing the form
</span></span></span><span style="display:flex;"><span><span style="color:#007f7f"></span>	ui := uni.NewFormLayout().
</span></span><span style="display:flex;"><span>		Add(<span style="color:#0ff;font-weight:bold">&#34;Name&#34;</span>, <span style="color:#0ff;font-weight:bold">&#34;Name&#34;</span>)
</span></span><span style="display:flex;"><span>
</span></span><span style="display:flex;"><span>	<span style="color:#007f7f">// Errors are a map of string -&gt; string
</span></span></span><span style="display:flex;"><span><span style="color:#007f7f"></span>	<span style="color:#007f7f">// with field names and error texts
</span></span></span><span style="display:flex;"><span><span style="color:#007f7f"></span>	errors := <span style="color:#fff;font-weight:bold">map</span>[<span style="color:#fff;font-weight:bold">string</span>]<span style="color:#fff;font-weight:bold">string</span>{<span style="color:#0ff;font-weight:bold">&#34;Name&#34;</span>: <span style="color:#0ff;font-weight:bold">&#34;Name can&#39;t be Unicornus&#34;</span>}
</span></span><span style="display:flex;"><span>
</span></span><span style="display:flex;"><span>	<span style="color:#007f7f">// Render form layout with data
</span></span></span><span style="display:flex;"><span><span style="color:#007f7f"></span>	<span style="color:#007f7f">// to html
</span></span></span><span style="display:flex;"><span><span style="color:#007f7f"></span>	html := ui.RenderFormWithErrors(d, errors)
</span></span></code></pre>

Results in

<img src="https://raw.githubusercontent.com/inkmi/unicornus/master/formexample.png" width="600">



