<template>
  <div id="app" class="container-fluid">
    <nav class="navbar navbar-expand-lg navbar-light bg-light">
      <a
        class="navbar-brand"
        href="#"
        style="font-family: 'Jost', sans-serif;"
      >Blood Pressure Tracker</a>
      <button
        class="navbar-toggler"
        type="button"
        data-toggle="collapse"
        data-target="#navbarNavAltMarkup"
        aria-controls="navbarNavAltMarkup"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarNavAltMarkup">
        <div class="navbar-nav">
          <router-link class="nav-item nav-link" to="/today">Home</router-link>
          <router-link class="nav-item nav-link" to="/history">History</router-link>
        </div>
      </div>
    </nav>
    <router-view></router-view>
  </div>
</template>
