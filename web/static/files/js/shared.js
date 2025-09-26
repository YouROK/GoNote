function showMessage(text, duration = 2000) {
    const msgBox = document.getElementById('msgBox');
    if (!msgBox) return;
    msgBox.textContent = text;
    msgBox.classList.add('show');
    if (msgBox.hideTimeout) clearTimeout(msgBox.hideTimeout);
    msgBox.hideTimeout = setTimeout(() => {
        msgBox.classList.remove('show');
    }, duration);
}

function generatePassword(length = 24) {
    const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
    let pass = "";
    for (let i = 0; i < length; i++) {
        pass += chars.charAt(Math.floor(Math.random() * chars.length));
    }
    return pass;
}

// Экспортируем в глобальную область (если не используешь ES6 modules)
window.GonoteUtils = { showMessage, generatePassword };