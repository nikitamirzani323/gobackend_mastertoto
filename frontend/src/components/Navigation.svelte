<script>
  import { link } from "svelte-spa-router";
  async function handleLogout() {
    localStorage.clear();
    window.href = "/";
  }
</script>

<header>
  <!-- Fixed navbar -->
  <nav
    class="navbar navbar-expand-md navbar-dark fixed-top bg-dark"
    style="background-color: #1f2937!important;"
  >
    <div class="container-fluid">
      <a class="navbar-brand" href="/">BACKEND</a>
      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarCollapse"
        aria-controls="navbarCollapse"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon" />
      </button>
      <div class="collapse navbar-collapse" id="navbarCollapse">
        <ul class="navbar-nav me-auto mb-2 mb-md-0">
          <li class="nav-item">
            <a class="nav-link active" aria-current="page" use:link href="/"
              >DASHBOARD</a
            >
          </li>
          <li class="nav-item dropdown">
            <a
              class="nav-link dropdown-toggle"
              id="navbarDropdown"
              role="button"
              data-bs-toggle="dropdown"
              aria-expanded="false"
            >
              COMPANY
            </a>
            <ul class="dropdown-menu" aria-labelledby="navbarDropdown">
              <li>
                <a class="dropdown-item" use:link href="/company">COMPANY</a>
              </li>
            </ul>
          </li>
          <li class="nav-item dropdown">
            <a
              class="nav-link dropdown-toggle"
              id="navbarDropdown"
              role="button"
              data-bs-toggle="dropdown"
              aria-expanded="false"
            >
              MASTER
            </a>
            <ul class="dropdown-menu" aria-labelledby="navbarDropdown">
              <li>
                <a class="dropdown-item" use:link href="/pasaran">PASARAN</a>
              </li>
            </ul>
          </li>
        </ul>
        <form class="d-flex">
          <button 
            on:click={() => {
              handleLogout();
            }}
            type="button" 
            class="btn btn-warning">
            LOGOUT
          </button>
        </form>
      </div>
    </div>
  </nav>
</header>