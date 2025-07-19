import { API } from "./services/API.js";
import { HomePage } from "./components/HomePage.js";
import './components/AnimatedLoading.js';
import './components/YoutubeEmbed.js';
import { MovieDetailsPage } from "./components/MovieDetailsPage.js";
import { Router } from "./services/Router.js";
import proxiedStore from "./services/Store.js";

window.addEventListener("DOMContentLoaded", () => {
    app.Router.init();
});

window.app = {
    Router,
    proxiedStore,
    showError: (message = "There was an error", goToHome=false) => {
        document.getElementById("alert-modal").showModal();
        document.querySelector("#alert-modal p").textContent = message;
        if (goToHome) app.Router.go("/");
    },
    closeError: () => {
        document.getElementById("alert-modal").close();
    },
    search: (event) => {
        event.preventDefault();
        const q = document.querySelector("input[type=search]").value;
        app.Router.go("/movies?q=" + q);
    },
    searchOrderChange: (order) => {
        const urlParams = new URLSearchParams(window.location.search);
        const q = urlParams.get("q");
        const genre = urlParams.get("genre") ?? "";
        app.Router.go(`/movies?q=${q}&order=${order}&genre=${genre}`);
    },
    searchFilterChange: (genre) => {
        const urlParams = new URLSearchParams(window.location.search);
        const q = urlParams.get("q");
        const order = urlParams.get("order") ?? "";
        app.Router.go(`/movies?q=${q}&order=${order}&genre=${genre}`);
    },
    register: async (event) => {
        event.preventDefault();
        const name = document.getElementById("register-name").value;
        const password = document.getElementById("register-password").value;
        const passwordConfirm = document.getElementById("register-password-confirm").value;
        const email = document.getElementById("register-email").value;

        const errors = [];
        if (name.length < 4) errors.push("Enter your complete name");
        if (password.length < 7) errors.push("Enter a password with at least 7 characters");
        if (email.length < 4) errors.push("Enter your complete email");
        if (passwordConfirm != password) errors.push("Password's don't match");

        if (errors.length == 0) {
            const response = await API.register(name, email, password)
            if (response.success) {
                app.proxiedStore.jwt = response.jwt;
                app.Router.go("/account/")
            } else {
                app.showError(response.message);
            }
        } else {
            app.showError(errors.join(". "))
        }
    },
    login: async (event) => {
        event.preventDefault();
        let errors = [];
        const email = document.getElementById("login-email").value;
        const password = document.getElementById("login-password").value;
 
        if (email.length < 8) errors.push("Enter your complete email");
        if (password.length < 6) errors.push("Enter a password with 6 characters");
        if (errors.length==0) {
            const response = await API.login(email, password);
            if (response.success) {
                app.proxiedStore.jwt = response.jwt;
                app.Router.go("/account/")
            } else {
                app.showError(response.message, false);
            }
        } else {
            app.showError(errors.join(". "), false);
        }
    },
    api: API,
}