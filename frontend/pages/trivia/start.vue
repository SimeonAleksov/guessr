<template>
  <div class="flex flex-col justify-center items-center" v-show="$auth.loggedIn">
    <h1 class="font-mono text-4xl dark:text-secondary mb-4">
      Create a game or join an existing one!
    </h1>
    <label class="font-mono">
      Code
    </label>
    <Input type="text" placeholder="g13gk1k2j" :model="code" @input="onInput" @change="onChange"/>
    <Button class="mt-4" @click="start">
      Start!
    </Button>
  </div>
</template>

<script>
import Input from "~/components/shared/Input";
import Button from "~/components/shared/Button";
import { actions } from "@/constants/trivia";

export default {
  name: "start",
  middleware: 'auth',
  components: {
    Button,
    Input,
  },
  data() {
    return {
      code: '',
      connection: null,
    }
  },
  mounted() {
    console.log(this.$auth.strategy.token.get().split(" ")[1])
  },
  methods: {
    start() {
      this.connection = new WebSocket("ws://localhost:8000/ws/")
      this.connection.onopen = () => {
        this.sendMessage({
          action: actions.START,
          topic: this.code,
          token: this.$auth.strategy.token.get().split(" ")[1],
        });
      };
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
  }
}
</script>
