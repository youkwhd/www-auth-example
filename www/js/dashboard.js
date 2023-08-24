(async () => {
    let data;

    try {
        const response = await fetch("http://localhost:3000/user", { 
            credentials: "include",
            method: "GET", 
        });

        data = await response.json();

        if (!data.success) {
            window.location.assign("/login.html");
            return;
        }
    } catch (err) {
        window.location.assign("/login.html");
        return;
    }

    /* Imagine this is modern front-end library 
     * like React (I'm lazy to use one)
     */
    const body = document.querySelector("body");

    let headlineEl = document.createElement("h1");
    headlineEl.innerText = "Dashboard";

    let divEl = document.createElement("div");
    divEl.className = "form-group";

    let horizontalLineEl = document.createElement("hr");
    let formEl = document.createElement("form");

    let labelEl = document.createElement("label");
    labelEl.innerText = "Username";

    let inputEl = document.createElement("input");
    inputEl.readOnly = true;
    inputEl.value = data.username;

    divEl.appendChild(labelEl);
    divEl.appendChild(inputEl);

    formEl.appendChild(divEl);

    body.appendChild(headlineEl);
    body.appendChild(horizontalLineEl);
    body.appendChild(formEl);
})();
