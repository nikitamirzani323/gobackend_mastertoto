<script>
	import Router from "svelte-spa-router";
	import { wrap } from "svelte-spa-router/wrap";
	import Navigation from "./components/Navigation.svelte";
	import Dashboard from "./pages/dashboard/Dashboard.svelte";
	import Login from "./pages/Login.svelte";
	import NotFound from "./pages/Notfound.svelte";
	let token = localStorage.getItem("token");
	let routes = "";
	let isNav = false;
	if (token == "" || token == null) {
		routes = {
			"/": wrap({
				component: Login,
			}),
			"*": NotFound,
		};
	} else {
		isNav = true;
		routes = {
			"/": wrap({
				component: Dashboard,
			}),
			"*": NotFound,
		};
	}
</script>

{#if isNav}
	<Navigation />
{/if}
<Router {routes} />