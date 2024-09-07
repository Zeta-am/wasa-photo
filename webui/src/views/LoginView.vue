<script>
export default {
    data() {
        return {
            username: "",
            userId: null,
            errorMsg: "",
            isLoading: false,
            usrnameEval: new RegExp('^[a-zA-Z0-9_$]{3,16}$')
        }
    },
    methods: {
        async doLogin() {
            this.isLoading = true;
            try {
                if (!this.usrnameEval.test(this.username)) throw "Invalid username. It must contain minimum 3 characters and maximum 16 characters "
                let response = await this.$axios.post('/session', {
                    usrname: this.username,
                });

                if (response.status == 200 || response.status == 201) {
                    this.userId = response.data;
                    localStorage.token = this.userId;
                    localStorage.username = this.username;
                    this.$setAuth();
                    this.$router.push({name: 'Home'});
                }
            } catch(e) {
                const {response} = e;
                if (response.status == 400) {
                    this.errorMsg = "Username required"
                }
            }
            this.isLoading = false
        },
    },
    mounted() {
        var currId = this.$getCurrentId();
        // Check if the user is already logged in
        if (currId) {
            this.$router.push({name: 'Home'});
        }
    }
}
</script>

<template>
    <div class="form-body">
        <div style="max-width: 100vh;" class="form-inner">
            <div class="mb-3">
                <h1>WASAPhoto Login</h1>
            </div>
			<div style="font-size: 1.3rem;">
				<div class="form-group mb-3">
					<label for="Username">Username</label>
					<input style="text-align: center;" v-model="username" id="Username" type="text" class="form-control" placeholder="Type username...">
				</div>
				<button style="font-size:1.2rem;" class="profile-buttons profile-buttons-primary" type="button" @click="doLogin">
					<svg style="margin-left: 2px;" class="feather"><use href="/feather-sprite-v4.29.0.svg#log-in"/></svg>
					Login
				</button>
			</div>
			<div v-if="errorMsg" >
				<hr>
				<ErrorMsg :msg="errorMsg"></ErrorMsg>
			</div>
        </div>
    </div>
</template>
<!-- <template>
    <div>
      <h1>Login</h1>
      <input v-model="username" placeholder="Username" />
      <button @click="doLogin">Login</button>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        username: "",
        errorMsg: "",
        isLoading: false,
      };
    },
    methods: {
      async doLogin() {
        this.isLoading = true;
        console.log("Login attempt for username:", this.username);
        this.isLoading = false;
      },
    },
  };
  </script> -->
  