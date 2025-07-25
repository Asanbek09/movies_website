import { routes } from "./Routes.js";

export const Router = {
    init: () =>{
        window.addEventListener("popstate", () => {
            Router.go(location.pathname, false);
        });
        
        document.querySelectorAll("a.navlink").forEach(a => {
            a.addEventListener("click", event => {
                event.preventDefault();
                const href = a.getAttribute("href");
                Router.go(href);
            })
        })
        // go to the initial route
        Router.go(location.pathname + location.search)
    },
    go: (route, addToHistory=true) => {
        if (addToHistory) {
            history.pushState(null, "", route)
        }
        let pageElement = null

        const routePath = route.includes('?') ? route.split("?")[0] : route;
        
        let needsLogin = false;

        for (const r of routes) {
            if (typeof r.path == "string" && r.path === routePath) {
                pageElement = new r.component();
                needsLogin = r.loggedIn === true
                break;
            } else if (r.path instanceof RegExp) {
                const match = r.path.exec(route);
                if (match) {
                    pageElement = new r.component();
                    const params = match.slice(1);
                    pageElement.params = params;
                    needsLogin = r.loggedIn === true
                    break;
                }
            }
        }

        if (pageElement) {
            if (needsLogin && app.proxiedStore.loggedIn == false) {
                app.Router.go("/account/login")
                return;
            }
        }

        if (pageElement == null) {
            pageElement = document.createElement("h1")
            pageElement.textContent = "Page not found"
        }
        // inserting the new page in the UI
        const oldPage = document.querySelector("main").firstElementChild;
        if (oldPage) oldPage.computedStyleMap.viewTransitionName = "old";
        pageElement.style.viewTransitionName = "new";

        function updatePage() {
            document.querySelector("main").innerHTML ="";
            document.querySelector("main").appendChild(pageElement);
        }

        if (!document.startViewTransition) {
            // we don't do a transition
            updatePage();
        } else {
            // we do a transition
            document.startViewTransition( () => {
                updatePage();
            });
        }
    }
}
