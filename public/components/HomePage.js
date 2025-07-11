export class HomePage extends HTMLElement { // <home-page>
    connectedCallback() {
        const template = document.getElementById("template-home");
        const content = template.content.cloneNode(true);
        this.appendChild(content);
    }
}

customElements.define("home-page", HomePage);