<template>
  <div class="container">
    <div class="switch">
      <a-button type="outline" shape="round" size="large">
        Register
      </a-button>
      <a-button type="outline" shape="round" size="large">
        Register
      </a-button>
      <a-button type="outline" shape="round" size="large">
        Register
      </a-button>
      <a-button type="outline" shape="round" size="large">
        Register
      </a-button>
      <a-button type="outline" shape="round" size="large">
        Register
      </a-button>
      <a-button type="outline" shape="round" size="large">
        Register
      </a-button>
    </div>
  </div>
</template>

<style scoped>
  .container {
    padding: 16px;
    display: flex;
    background-color: var(--success-1);
  }

  .switch {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    width: 100%;
  }

  .switch > * {
    margin: 0 40px;
  }
</style>
