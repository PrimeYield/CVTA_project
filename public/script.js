    function login() {
      const username = document.getElementById("username").value;
      const password = document.getElementById("password").value;

      if (username === "1234" && password === "1234") {
        document.getElementById("login").classList.add("hidden");
        document.getElementById("main").classList.remove("hidden");
      } else {
        document.getElementById("login-error").innerText = "帳號或密碼錯誤";
      }
    }


    
    function calculate() {
      const height = parseFloat(document.getElementById("height").value);
      const weight = parseFloat(document.getElementById("weight").value);
      const exercise = document.getElementById("exercise").value;
      const count = parseInt(document.getElementById("count").value);

      if (!height || !weight || !count) {
        document.getElementById("result").innerText = "請填寫完整資料";
        return;
      }

      // 假設每次運動消耗熱量（大致估算）
      const calorieRates = {
        situp: 0.5,
        pushup: 0.8,
        squat: 0.7,
      };

      // 假設標準次數（可以依年齡或性別再調整）
      const standard = {
        situp: 30,
        pushup: 20,
        squat: 25,
      };

      const isPass = count >= standard[exercise];
      const calories = (calorieRates[exercise] * weight * count / 60).toFixed(2);

      document.getElementById("result").innerText =
        `你做了 ${count} 次${getExerciseName(exercise)}。\n` +
        `符合標準？${isPass ? "是 ✅" : "否 ❌"}\n` +
        `大約消耗熱量：${calories} 大卡`;
    }

    function getExerciseName(code) {
      return {
        situp: "仰臥起坐",
        pushup: "伏地挺身",
        squat: "深蹲",
      }[code];
    }