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
                let response = await this.$axios.post('/users', {
                    username: this.username,
                });

                if (response.status == 200 || response.status == 201) {
                    this.userId = response.data;
                    localStorage.token = this.userId;
                    localStorage.username = this.username;
                    this.$setAuth();
                    this.$router.push({name: 'Home'});
                }
            } catch(e) {
                console.log(e)
                const response = e.response;
                if (response) {
                    if (response.status == 400) {
                        this.errorMsg = "Username required"
                    } else {
                        this.errorMsg = "An error occured " + (response.data?.message || "Unknown error");
                    }
                } else {
                    this.errorMsg = "An unexpected error occured";
                }
            }
            this.isLoading = false
        },
    },
    mounted() {
        var currId = this.$utils.getCurrentId();
        // Check if the user is already logged in
        if (currId) {
            this.$router.push({name: 'Home'});
        }
    }
}
</script>

<template>
    <!-- <div class="form-body">
        <div style="max-width: 100vh;" class="form-inner">
            <div class="mb-3">
                <h1>WASAPhoto Login</h1>
            </div>
			<div style="font-size: 1.3rem;">
				<div class="form-group mb-3">
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
    </div> -->
    <div class="form-body">
        <div class="form-inner">
            <!-- Title or image -->
            <div class="mb-3 text-center">
                <!-- Title -->
                <h1 class="styled-title"> WasaPhoto </h1>
                
                <!-- Image -->
                <!-- <img src="" alt="WasaPhotoIcon" class="login-logo"> -->
            </div>

            <!-- Input text field and login button -->
            <div class="form-content text-center">
                <div class="form-group mb-3">
                    <input 
                        v-model="username"
                        id="Username"
                        type="text"
                        class="form-control text-center"
                        placeholder="Type username..."
                    >
                </div>
                <button class="styled-button" type="button" @click="doLogin">
                    Log in
                </button>
            </div>
            <!-- Error message -->
            <div v-if="errorMsg">
                <hr>
                <ErrorMsg :msg="errorMsg"></ErrorMsg>
            </div>
        </div>
    </div>
    
</template>

<style>
.form-body {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 100vh; /* Per centrare verticalmente */
}

.form-inner {
    max-width: 400px; /* Imposta una larghezza massima */
    width: 100%; /* Imposta la larghezza al 100% */
    padding: 20px;
    background-color: white; /* Colore di sfondo (se necessario) */
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1); /* Leggera ombra per effetto 3D */
    border-radius: 8px; /* Angoli arrotondati */
}

.text-center {
    text-align: center; /* Centra il testo */
}

.login-logo {
    max-width: 100%; /* Assicura che l'immagine sia adattata correttamente */
    height: auto; /* Mantiene il rapporto di aspetto dell'immagine */
    margin-bottom: 20px; /* Spazio sotto l'immagine */
}

.form-content {
    display: flex;
    flex-direction: column; /* Disposizione verticale */
    align-items: center; /* Allineamento orizzontale centrale */
}

.form-control {
    text-align: center; /* Centra il testo all'interno del campo di input */
    width: 100%; /* Utilizza l'intera larghezza */
    padding: 10px; /* Spazio all'interno del campo */
    margin-bottom: 15px; /* Spazio sotto il campo */
    font-size: 1.2rem; /* Dimensione del testo */
}

.profile-buttons {
    font-size: 1.2rem; /* Dimensione del testo per il bottone */
    padding: 10px 20px; /* Spazio interno del bottone */
    cursor: pointer; /* Puntatore quando si passa sopra il bottone */
}

.styled-title {
    font-family: 'Arial', sans-serif; /* Font semplice ma elegante */
    font-size: 2.5rem; /* Dimensione del titolo più grande */
    color: #333; /* Colore del testo */
    margin-bottom: 30px; /* Spazio inferiore */
    font-weight: bold; /* Grassetto */
    text-transform: uppercase; /* Maiuscolo */
    letter-spacing: 2px; /* Spaziatura tra le lettere */
    position: relative;
}

.styled-title::after {
    content: ''; /* Aggiunge una linea decorativa sotto il titolo */
    width: 200px;
    height: 5px;
    background-color: #4ba8b9; /* Colore accentato */
    display: block;
    margin: 10px auto 0; /* Allineamento e spazio */
    border-radius: 5px;
}

.styled-button {
    background-color: #4ba8b9; /* Colore di sfondo */
    color: #fff; /* Colore del testo */
    font-size: 1.2rem;
    font-weight: bold;
    padding: 10px 20px;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    display: flex; /* Per centrare l'icona e il testo */
    align-items: center; /* Centra verticalmente */
    transition: background-color 0.3s ease, transform 0.3s ease; /* Animazione di hover */
}

.styled-button:hover {
    background-color: #4ba8b9; /* Colore di hover */
    transform: translateY(-2px); /* Effetto sollevamento */
}

.styled-button:active {
    background-color: #4ba8b9; /* Colore quando cliccato */
    transform: translateY(0); /* Rimuove l'effetto di sollevamento */
}
</style>