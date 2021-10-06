<script>
    let username = "";
    let password = "";
    let msg;
    let flag;

    async function handleSubmit() {
        flag = true;
        msg = "";
        if (username == "") {
            flag = false;
            msg += "The username is required\n";
        }
        if (password == "") {
            flag = false;
            msg += "The Password is required\n";
        }
        if (flag) {
            const res = await fetch("/api/login", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    username: username,
                    password: password,
                }),
            });
            const json = await res.json();
            if (json.status == 400) {
                alert(json.message);
                username = "";
                password = "";
            } else {
                localStorage.setItem("token", json.token);
                localStorage.setItem("master", username);
                window.location.href = "/";
            }
        } else {
            alert(msg);
        }
    }
</script>

<div class="container" style="margin-top:70px">
    <div class="row">
        <center>
            <div class="col-sm-3">
                <div class="card" style="border-radius:0px;margin-top:10px;">
                    <div class="card-header" style="background-color: #1f2937;">
                        <span
                            style="color:white;font-weight: bold;font-size: 15px;"
                            >LOGIN - BACKEND</span
                        >
                    </div>
                    <div class="card-body">
                        <div class="mb-3">
                            <input
                                bind:value={username}
                                type="text"
                                class="form-control"
                                placeholder="Username"
                                aria-label="Username"
                            />
                        </div>
                        <div class="mb-3">
                            <input
                                bind:value={password}
                                type="password"
                                class="form-control"
                                placeholder="Password"
                                aria-label="Password"
                            />
                        </div>
                        <div class="mb-3">
                            <div class="d-grid gap-1">
                                <button 
                                    on:click={() => {
                                        handleSubmit();
                                    }}
                                    type="button" 
                                    class="btn btn-primary">
                                    SUBMIT
                                </button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </center>
    </div>
</div>