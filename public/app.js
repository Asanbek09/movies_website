import { API } from "./services/API.js";
import { HomePage } from "./components/HomePage.js";
import './components/AnimatedLoading.js';
import './components/YoutubeEmbed.js';
import { MovieDetailsPage } from "./components/MovieDetailsPage.js";
import { Router } from "./services/Router.js";

window.addEventListener("DOMContentLoaded", () => {
    app.Router.init();
});

window.app = {
    API,
    Router,
    search: (event) => {
        event.preventDefault();
        const q = document.querySelector("input[type=search]").value;
    },
}