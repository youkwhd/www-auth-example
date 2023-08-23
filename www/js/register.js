window.onload = () => {
    const form = document.querySelector("form");

    form.addEventListener("submit", async (e) => {
        e.preventDefault();

        const username = e.target.querySelector('[name="input--username"]').value;
        const password = e.target.querySelector('[name="input--password"]').value;

        const credentials = {
            username, 
            password
        };

        const response = await fetch("http://localhost:3000/register", { 
            method: "POST", 
            credentials: "include",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify(credentials),
        });

        const data = await response.json();

        if (data.success) {
            window.location.assign("/dashboard.html");
            return;
        }
    })
}
