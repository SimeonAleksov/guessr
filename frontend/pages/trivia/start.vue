<template>
  <div class="flex flex-col justify-center items-center" v-show="$auth.loggedIn">
    <div class="flex flex-col justify-center items-center" v-if="action === null">
    <h1 class="font-mono text-4xl dark:text-secondary mb-4">
      Create a game or join an existing one!
    </h1>
      <Grid :albums="categories" @click="onImageClick" class="ease-in duration-500"/>
    <label class="font-mono mb-2">
      Code
    </label>
    <h1 v-show="create" class="font-mono text-4xl dark:text-primary">
      {{ code }}
    </h1>
    <Input v-show="!create" type="text" :placeholder="code" :model="code" @input="onInput" @change="onChange"/>
    <div class="mb-4 mt-4 flex flex-col justify-content items-center max-h-xl">
      <h2 class="text-center text-secondary font-semibold mt-4 mb-4" v-if="create">
        Toggle to join a game!
      </h2>
      <h2 class="text-center text-secondary font-semibold mt-4 mb-4" v-else>
        Toggle to create a game!
      </h2>
      <Toggle :value="create" @input="onToggleInput"/>
    </div>
    <Button class="mt-4" @click="start">
      {{ create ? 'Create a game!' : 'Join a game!'}}
    </Button>
      <h1 class="mt-8 font-mono text-2xl dark:text-primary mb-4">
        Currently selected: {{ selectedCategoryName }}
      </h1>
    </div>
    <div class="mt-12 flex space-x-4 z-1">
      <div v-for="user in users" :key="user.username" class="dashboard__player dark:bg-dark text-primary px-4 uppercase flex flex-col justify-center">
        <h1 class="text-center text-secondary font-semibold">0</h1>
        <img class="w-24 h-24" :src="`https://robohash.org/${user.username}`" alt=""/>
        <h1 class="text-center">{{ user.username }}</h1>
      </div>
    </div>
    <div v-if="action === actions?.CREATED" class="flex flex-col justify-center items-center">
      <Button class="mt-4" @click="startGame">
        Start!
      </Button>
    </div>
    <div v-if="action === actions?.JOINED">
      <h1 class="mt-8 font-mono text-2xl dark:text-secondary mb-4">
        Waiting for the host to start the game!
      </h1>
    </div>
    <div>
      <h1 v-show="countdown" class="mt-8 font-mono text-2xl dark:text-primary mb-4">
        Game starting in: {{ countdown }}
      </h1>
    </div>
  </div>
</template>

<script>
import Input from "~/components/shared/Input";
import Button from "~/components/shared/Button";
import {actions} from "@/constants/trivia";
import {messages} from "@/constants/messages";
import {images} from "@/constants/pfp";
import Toggle from "@/components/shared/Toggle";

export default {
  name: "start",
  middleware: 'auth',
  components: {
    Toggle,
    Button,
    Input,
  },
  data() {
    return {
      create: true,
      code: '',
      connection: null,
      users: null,
      pfpImages: images,
      action: null,
      actions: {},
      categories: [],
      selectedCategory: null,
      selectedCategoryName: null,
      ticker: null,
    }
  },
  computed: {
   countdown() {
     return this.ticker;
   }
  },
  mounted() {
    this.pfpImages = images;
    this.actions = messages;
    this.$axios.get('http://localhost:8000/trivia/categories/', {
      headers: {
        Authorization: this.$auth.strategy.token.get(),
      }
    })
      .then(resp => {
        this.categories = resp.data.data;
      })
      .catch(() => {})
    this.code = (Math.random() + 1).toString(36).substring(2).toUpperCase()

  },
  methods: {
    start() {
      this.connection = new WebSocket("ws://localhost:8000/ws/")
      this.connection.onopen = () => {
        this.sendMessage({
          action: actions.CREATE_OR_JOIN,
          topic: this.code,
          token: this.$auth.strategy.token.get().split(" ")[1],
          data: {
            triviaID: parseInt(this.selectedCategory),
            code: this.code,
          }
        });
      };
      this.connection.onmessage = (message) => {
        try {
          let parsedData = JSON.parse(message.data)
          console.log(parsedData)
          if (parsedData.type === messages.CURRENT_USERS) {
            this.users = parsedData.users
          } else if (parsedData.type === messages.CREATED) {
            this.action = messages.CREATED
          } else if (parsedData.type === messages.JOINED) {
            this.action = messages.JOINED
          } else if (parsedData.action === messages.TICKER) {
            this.ticker = parseInt(parsedData.data)
          }
        } catch (e) {
          console.log(e)
        }
      }
    },
    startGame() {
      this.sendMessage({
        action: actions.START,
        topic: this.code,
        token: this.$auth.strategy.token.get().split(" ")[1],
      })
    },
    onInput(value) {
      this.code = value;
    },
    onChange(value) {
      this.code = value;
    },
    sendMessage(message) {
      this.connection.send(JSON.stringify(message));
    },
    onImageClick(e) {
      this.selectedCategory = e;
      this.selectedCategoryName = this.categories.find(x => x.ID === parseInt(this.selectedCategory))?.Name
    },
    onToggleInput(value) {
      this.create = value;
    },
  },
  beforeDestroy() {
    this.connection.onclose = () => {
      this.sendMessage({
        action: actions.UNSUBSCRIBE,
        topic: this.code,
        token: this.$auth.strategy.token.get().split(" ")[1],
      });
    }
  }
}
</script>
