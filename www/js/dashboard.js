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

    let descriptionEl = document.createElement("p");
    descriptionEl.innerText = "This is a user only page, you can see this if you are authenticated.";

    let horizontalLineEl = document.createElement("hr");
    let formEl = document.createElement("form");

    let labelEl = document.createElement("label");
    labelEl.innerText = "Username ";

    let readOnlyEl = document.createElement("span");
    readOnlyEl.className = "read-only-msg";
    readOnlyEl.innerText = "(Read Only)";

    let inputEl = document.createElement("input");
    inputEl.readOnly = true;
    inputEl.value = data.username;

    let logoutButtonEl = document.createElement("button");
    logoutButtonEl.innerText = "Log out";

    logoutButtonEl.addEventListener("click", async () => {
        const response = await fetch("http://localhost:3000/logout", { 
            credentials: "include",
            method: "GET", 
        });

        const data = await response.json();

        if (data.success) {
            window.location.assign("/index.html");
            return;
        }
    });

    labelEl.appendChild(readOnlyEl);

    divEl.appendChild(labelEl);
    divEl.appendChild(inputEl);

    formEl.appendChild(divEl);

    body.appendChild(headlineEl);
    body.appendChild(horizontalLineEl);
    body.appendChild(descriptionEl);
    body.appendChild(formEl);
    body.appendChild(logoutButtonEl);
})();
