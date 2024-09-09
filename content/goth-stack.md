---
name: "Introducing the GOTH Stack"
date: "January 08 2024"
tags: ["article","golang"]
---
# Enter the GOTH Stack!
The GOTH stack is something I've been trying to get to for a while now. It's not a specific repository with a fancy command that can scaffold a project for you. It's more like a set of pillars for building excellent, pleasant, full-stack web applications.

# The first pillar: Go
Go is something I learned to love later on in my career. I was mainly writing JavaScript, building on serverless platforms and growing frustrated at the performance and limitations. Go changed all of that.

What makes Go good?:
- Static types
- Incredibly easy concurrency
- Errors as values
- Incredible runtime and build time performance
- Tiny memory footprint

The tl;dr is that it is challenging to write Go code that is _not_ performant.

# The second pillar: Templates... well... Go templates
Go templates surprised me, to be completely honest. They offer just enough to get me going and perform exceptionally well. Sure, it's not as simple as a basic JSX file in Next.js since you need to make a route handler, but it works pretty well and supports basic control flow. I am interested in looking into alternatives such as TEMPL (which reads much like JSX for Go), but I need to find a real reason to move from the standard library here.

Here is an example of a route handler passing a slice over to a template for rendering:
```go
package pages

import (
    "HTML/template"

    "github.com/labstack/echo/v4"
    "goth.stack/lib"
)

type HomeProps struct {
    Socials      []lib.IconLink
    Tech         []lib.IconLink
    ContractLink string
    ResumeURL    string
    SupportLink  string
}

func Home(c echo.Context) error {
    socials := []lib.IconLink{
        {
            Name: "Email",
            Href: "mailto:example@site.com",
            Icon: template.HTML(`<svg xmlns="http://www.w3.org/2000/svg" height="32" width="32" viewBox="0 0 512 512"><!--!Font Awesome Free 6.5.1 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2023 Fonticons, Inc.--><path d="M48 64C21.5 64 0 85.5 0 112c0 15.1 7.1 29.3 19.2 38.4L236.8 313.6c11.4 8.5 27 8.5 38.4 0L492.8 150.4c12.1-9.1 19.2-23.3 19.2-38.4c0-26.5-21.5-48-48-48H48zM0 176V384c0 35.3 28.7 64 64 64H448c35.3 0 64-28.7 64-64V176L294.4 339.2c-22.8 17.1-54 17.1-76.8 0L0 176z"/></svg>`),
        },
    }

    props := HomeProps{
        Socials:      socials,
        Tech:         tech,
        ContractLink: "mailto:example@site.com",
        ResumeURL:    "https://srv.goth.stack/Atridad_Lahiji_Resume.pdf",
        SupportLink: "https://donate.stripe.com/8wMeVF25c78L0V2288",
    }

    // Specify the partials used by this page
    partials := []string{"header", "navitems"}

    // Render the template
    return lib.RenderTemplate(c.Response().Writer, "base", partials, props)
}
```
As you can see, it really isn't that bad! It also comes with many of the benefits of Go and the flexibility of components!

# The third pillar: HTMX
So, up to this point, you may have been thinking: "Gee Atri... you can't do anything reactive here". Before HTMX, you would have been right. HTMX offers a more backend-centric developer a way to build complex reactivity to their front end through basic templating languages. It is one file you import in your template, and it enables anything from basic HTML swapping to WebSocket and Server-Sent Event support. It is really, really powerful and worth looking at all together.

With Go managing route handlers and API routes, the template language running the UI, and HTMX governing interactivity on the front end, you can effectively write a fully dynamic full-stack application without writing a line of JavaScript code. It even runs quite a bit faster than some of the JS world's frameworks (Astro, for instance). It is actually what powers this site right now! The fundamentals are essential here and come together for a clean and enjoyable developer experience. I encourage everyone in the JS world to give it a shot! Perhaps it is not your thing, and that's okay! But you might also just fall in love with it!

# Thanks!
If you found this helpful, please let me know by email at [me@atri.dad](mailto:me@atri.dad). Until next time!