<template>
    <div class="home-page">
      <!-- Website Name -->
      <h1>WhichRig</h1>
  
      <!-- Site Description -->
      <p>
        WhichRig helps you compare used computers and computer parts. Whether you're building a new system, upgrading, or looking for the best deals, we've got you covered.
      </p>
  
      <!-- Links Section -->
      <div class="links">
        <ul>
          <li>
            <router-link to="/comparestuff">Enter your computer specs or give a link to get a rating or compare multiple computers</router-link>
          </li>
          <li>
            <router-link to="/top-components">See the top components for price/performance per category</router-link>
          </li>
          <li>
            <router-link to="/best-deals">See the best deals available by systems or by type of component</router-link>
          </li>
          <li>
            <router-link to="/saved-builds">See your saved builds or saved listings for computers or components</router-link>
          </li>
        </ul>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: 'HomePage'
  }
  </script>
  
  <style scoped>
  .home-page {
    text-align: center;
    margin: 20px;
  }
  
  h1 {
    font-size: 3rem;
    color: #42b983;
  }
  
  p {
    font-size: 1.5rem;
    margin-bottom: 40px;
  }
  
  .links ul {
    list-style-type: none;
    padding: 0;
  }
  
  .links li {
    margin: 15px 0;
  }
  
  .links a {
    text-decoration: none;
    color: #3498db;
    font-weight: bold;
  }
  
  .links a:hover {
    color: #2980b9;
  }
  </style>
  