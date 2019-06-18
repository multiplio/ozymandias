import React from 'react';

const Login = () =>
  <div>
    <h2>Authenticate with github to create an access token</h2>
    <form action="/auth/auth" method="post">
      <input type="text" title="username" placeholder="username" />
      <input type="password" title="username" placeholder="password" />
      <button type="submit" class="btn">Login</button>
    </form>
  </div>

export default Login
