import React from 'react';

const Login = () =>
  <div>
    <h2>Authenticate with github to create an access token</h2>
    <form action="/auth/auth" method="post">
      <input type="text" name="username" placeholder="username" />
      <input type="password" name="password" placeholder="password" />
      <button type="submit" value="submit">Login</button>
    </form>
  </div>

export default Login
