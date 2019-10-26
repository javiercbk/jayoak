<script>
	import 'bulma/css/bulma.css';
	export let segment;
</script>

<nav class="navbar" role="navigation" aria-label="main navigation">
  <div class="navbar-brand">
    <a class="navbar-item" href="/">
      <object data="/nut.svg" type="image/svg+xml" width="112" height="28" aria-label="nut"></object>
    </a>

    <a href="" role="button" class="navbar-burger burger" aria-label="menu" aria-expanded="false" data-target="navbarBasicExample">
      <span aria-hidden="true"></span>
      <span aria-hidden="true"></span>
      <span aria-hidden="true"></span>
    </a>
  </div>

  <div id="navbarBasicExample" class="navbar-menu">
    <div class="navbar-start">
      <a class="navbar-item" href="/">
        Home
      </a>

      <a class="navbar-item" href="/instruments">
        Upload Instruments
      </a>

      <div class="navbar-item has-dropdown is-hoverable">
        <a class="navbar-link">
          User
        </a>

        <div class="navbar-dropdown">
          <a class="navbar-item" href="/users/me/settings">
            Settings
          </a>
          <hr class="navbar-divider">
          <a class="navbar-item">
            Logout
          </a>
        </div>
      </div>
    </div>

  </div>
</nav>
