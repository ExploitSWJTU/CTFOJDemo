-- SWJTU CTF OJ - Challenge Data Seed
-- 从前端 mock 数据生成的 PostgreSQL INSERT 语句
-- 生成时间：2026-02-25

-- 清空现有数据（可选）
-- TRUNCATE TABLE challenges RESTART IDENTITY CASCADE;

-- 插入 20 道题目
INSERT INTO challenges (title, description, category, points, difficulty, type, flag, solved_count, created_at, updated_at) VALUES
-- Web 分类
(1, 'Easy SQL Injection', '
### Description

This is a basic SQL injection challenge. Your goal is to bypass the login page.

```sql
SELECT * FROM users WHERE username = ''$username'' AND password = ''$password''
```

### Goal

Find the flag in the database.
', 'Web', 100, 'Easy', 'static', 'flag{sql_b4sic_inj3cti0n}', 120, NOW(), NOW()),

(2, 'Advanced XSS', '
### Description

Bypass the Content Security Policy to execute your JavaScript code.

### Goal

Steal the admin cookie or trigger an alert.
', 'Web', 300, 'Hard', 'static', 'flag{xss_csp_byp4ss}', 15, NOW(), NOW()),

(3, 'Prototype Pollution', '
### Description

Exploit prototype pollution vulnerability in a Node.js application.

```javascript
function merge(obj, source) {
  for (let key in source) {
    obj[key] = source[key];
  }
}
```

### Goal

Pollute the Object prototype to bypass authentication.
', 'Web', 250, 'Medium', 'static', 'flag{pr0t0typ3_p0lluti0n}', 30, NOW(), NOW()),

(4, 'JWT Secret Brute Force', '
### Description

Find the weak secret used to sign the JWT tokens.

### Goal

Forge an admin JWT token.
', 'Web', 150, 'Easy', 'static', 'flag{jwt_w34k_s3cr3t}', 95, NOW(), NOW()),

-- Pwn 分类
(5, 'Buffer Overflow Level 1', '
### Description

A simple buffer overflow vulnerability. Can you overwrite the return address?

```c
void vulnerable_function(char *input) {
    char buffer[64];
    strcpy(buffer, input);
}

int main(int argc, char **argv) {
    vulnerable_function(argv[1]);
    return 0;
}
```

### Goal

Get shell access by overwriting the return address.
', 'Pwn', 200, 'Medium', 'static', 'flag{buf_0v3rfl0w_b4sic}', 45, NOW(), NOW()),

(6, 'Kernel Exploit', '
### Description

Exploit a race condition vulnerability in the kernel module.

### Goal

Gain root privileges through TOCTOU exploit.
', 'Pwn', 500, 'Hard', 'static', 'flag{k3rn3l_r4c3_c0nditi0n}', 5, NOW(), NOW()),

(7, 'Heap Spraying', '
### Description

Master the art of heap spraying to control memory layout.

### Goal

Achieve arbitrary code execution through heap exploitation.
', 'Pwn', 450, 'Hard', 'static', 'flag{h34p_spr4y_t3ch}', 12, NOW(), NOW()),

(8, 'VM Escape', '
### Description

Escape from the custom VM sandbox.

### Goal

Break out of the virtual machine and read /flag.
', 'Pwn', 600, 'Hard', 'static', 'flag{vm_3sc4p3_m4st3r}', 3, NOW(), NOW()),

-- Crypto 分类
(9, 'RSA 101', '
### Description

Decrypt the message using the given N and e.

```
N = 123456789012345678901234567890
e = 65537
c = 987654321098765432109876543210
```

### Goal

Factor N and recover the private key to decrypt the flag.
', 'Crypto', 150, 'Easy', 'static', 'flag{rs4_f4ct0ring}', 80, NOW(), NOW()),

(10, 'Elliptic Curve Cryptography', '
### Description

Solve the discrete logarithm problem on an elliptic curve.

### Goal

Find the private key from the public key.
', 'Crypto', 400, 'Hard', 'static', 'flag{3cc_dlp_s0lv3r}', 18, NOW(), NOW()),

(11, 'Power Analysis Side Channel', '
### Description

Recover the encryption key using power consumption traces.

### Goal

Perform Differential Power Analysis (DPA) attack.
', 'Crypto', 450, 'Hard', 'static', 'flag{sid3_ch4nn3l_dpa}', 9, NOW(), NOW()),

-- Misc 分类
(12, 'Hidden Image', '
### Description

Find the hidden flag in the image metadata.

### Goal

Extract EXIF data and find the hidden message.
', 'Misc', 100, 'Easy', 'attachment', 'flag{m3t4d4t4_hidd3n}', 200, NOW(), NOW()),

-- Reverse 分类
(13, 'Easy Reverse', '
### Description

Find the password hidden in the binary.

### Goal

Reverse engineer the binary to find the correct input.
', 'Reverse', 100, 'Easy', 'static', 'flag{r3v3rs3_3ngin33ring}', 150, NOW(), NOW()),

(14, 'Logic Bomb', '
### Description

Defuse the logic bomb before it "explodes".

### Goal

Find and neutralize the trigger condition.
', 'Reverse', 300, 'Medium', 'static', 'flag{l0gic_b0mb_diffus3d}', 25, NOW(), NOW()),

-- Mobile 分类
(15, 'Android Backup', '
### Description

Extract data from the Android backup file.

### Goal

Recover deleted messages and find the flag.
', 'Mobile', 200, 'Medium', 'attachment', 'flag{4ndr0id_b4ckup}', 60, NOW(), NOW()),

(16, 'iOS App Static Analysis', '
### Description

Find vulnerabilities in the compiled iOS binary.

### Goal

Analyze the IPA file and find hardcoded secrets.
', 'Mobile', 350, 'Hard', 'attachment', 'flag{i0s_st4tic_4n4lysis}', 14, NOW(), NOW()),

-- Blockchain 分类
(17, 'Smart Contract Vulnerability', '
### Description

Exploit the reentrancy bug in the smart contract.

```solidity
function withdraw(uint _amount) public {
    require(balances[msg.sender] >= _amount);
    msg.sender.call.value(_amount)();
    balances[msg.sender] -= _amount;
}
```

### Goal

Drain all funds from the contract.
', 'Blockchain', 400, 'Hard', 'static', 'flag{r33ntr4ncy_4tt4ck}', 20, NOW(), NOW()),

(18, 'DeFi Flash Loan Attack', '
### Description

Simulate a flash loan attack on a DeFi protocol.

### Goal

Manipulate the price oracle to profit from arbitrage.
', 'Blockchain', 500, 'Hard', 'static', 'flag{fl4sh_l04n_4tt4ck}', 8, NOW(), NOW()),

-- AI 分类
(19, 'Adversarial Example', '
### Description

Create an image that fools the neural network.

### Goal

Generate an adversarial example that is classified incorrectly.
', 'AI', 350, 'Medium', 'attachment', 'flag{4dv3rs4ri4l_m4g}', 10, NOW(), NOW()),

(20, 'Deepfake Detection', '
### Description

Identify if the given video is a deepfake.

### Goal

Analyze the video and determine its authenticity.
', 'AI', 200, 'Easy', 'attachment', 'flag{d33pf4k3_d3t3ct3d}', 50, NOW(), NOW());

-- 验证插入结果
SELECT 
    category, 
    COUNT(*) as total,
    SUM(solved_count) as total_solves,
    AVG(points) as avg_points
FROM challenges 
GROUP BY category 
ORDER BY total DESC;

-- 按难度统计
SELECT 
    difficulty,
    COUNT(*) as total,
    SUM(solved_count) as total_solves
FROM challenges 
GROUP BY difficulty 
ORDER BY 
    CASE difficulty 
        WHEN ''Easy'' THEN 1 
        WHEN ''Medium'' THEN 2 
        WHEN ''Hard'' THEN 3 
    END;
