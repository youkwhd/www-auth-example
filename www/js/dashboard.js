window.onload = async () => {
    const response = await fetch("http://localhost:3000/user", { 
        credentials: "include",
        method: "GET", 
    });

    const data = await response.json();

    if (!data.success) {
        window.location.assign("/register.html");
        return;
    }

    /* Imagine this is modern front-end library 
     * like React (I'm lazy to use one)
     */
    const body = document.querySelector("body");

    const headlineEl = document.createElement("h1");
    headlineEl.innerText = "Dashboard";

    const usernameEl = document.createElement("p");
    usernameEl.innerText = "username: ";

    const boldEl = document.createElement("strong");
    boldEl.innerText = data.username;

    usernameEl.appendChild(boldEl);

    body.appendChild(headlineEl);
    body.appendChild(usernameEl);
}
