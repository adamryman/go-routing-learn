# Learn routing

Just following along with this blog post: (http://www.alexedwards.net/blog/a-recap-of-request-handling)

## Conclusion

After this exercise and reading a little I have decided not to use net/http

From this link: (http://stackoverflow.com/a/30063908)
>There are two downsides to the builtin mux:
>
>If you need info from the url (for example id in /users/:id) you have to do it manually:
```
http.HandleFunc("/users/", func(res http.ResponseWriter, req *http.Request) {
    >id := req.URL.Path.SplitN("/", 3)[2]
})
```
>Which is cumbersome.
>
>The default server mux is not the fastest.

This mux looks good to me (https://github.com/julienschmidt/httprouter) going to go with that for ambition.
