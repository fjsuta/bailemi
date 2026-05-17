/**
 * Bailemi Music Platform - Professional SVG Icon Library
 * 百米乐音乐平台 - 专业 SVG 矢量图标库
 */

export const Icons = {
  // 主 Logo 图标
  logo: `
    <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
      <defs>
        <linearGradient id="logoGrad" x1="0%" y1="0%" x2="100%" y2="100%">
          <stop offset="0%" style="stop-color:#8B5CF6;stop-opacity:1" />
          <stop offset="100%" style="stop-color:#3B82F6;stop-opacity:1" />
        </linearGradient>
      </defs>
      <path d="M12 3v10.55c-.59-.34-1.27-.55-2-.55-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4V7h4V3h-6z" fill="url(#logoGrad)"/>
    </svg>
  `,

  // 播放控制图标
  play: `
    <svg viewBox="0 0 24 24" fill="currentColor" xmlns="http://www.w3.org/2000/svg">
      <polygon points="5,3 19,12 5,21 5,3"/>
    </svg>
  `,

  pause: `
    <svg viewBox="0 0 24 24" fill="currentColor" xmlns="http://www.w3.org/2000/svg">
      <rect x="6" y="4" width="4" height="16"/>
      <rect x="14" y="4" width="4" height="16"/>
    </svg>
  `,

  prev: `
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" xmlns="http://www.w3.org/2000/svg">
      <polygon points="19,4 9,12 19,20 19,4"/>
      <polygon points="5,4 15,12 5,20 5,4"/>
    </svg>
  `,

  next: `
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" xmlns="http://www.w3.org/2000/svg">
      <polygon points="5,4 15,12 5,20 5,4"/>
      <polygon points="19,4 9,12 19,20 19,4"/>
    </svg>
  `,

  // 随机播放
  shuffle: `
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" xmlns="http://www.w3.org/2000/svg">
      <path d="M16,3h5v5"/>
      <path d="M4,20 21,3"/>
      <path d="M21,16v5h-5"/>
      <path d="M15,15l6,6"/>
      <path d="M4,4l5,5"/>
    </svg>
  `,

  // 循环播放
  repeat: `
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" xmlns="http://www.w3.org/2000/svg">
      <path d="M21,12a9,9 0 0,0 -9,-9a9.75,9.75 0 0,0 -6.74,2.74L3,8"/>
      <path d="M3,16a9,9 0 0,0 9,9a9.75,9.75 0 0,0 6.74,-2.74L21,16"/>
    </svg>
  `,

  repeatOne: `
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" xmlns="http://www.w3.org/2000/svg">
      <path d="M21,12a9,9 0 0,0 -9,-9a9.75,9.75 0 0,0 -6.74,2.74L3,8"/>
      <path d="M3,16a9,9 0 0,0 9,9a9.75,9.75 0 0,0 6.74,-2.74L21,16"/>
      <circle cx="12" cy="12" r="3"/>
    </svg>
  `,

  // 音量控制
  volume: `
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" xmlns="http://www.w3.org/2000/svg">
      <polygon points="11,5 6,9 2,9 2,15 6,15 11,19 11,5"/>
      <path d="M15.54,8.46a5,5 0 0,1 0,7.07"/>
      <path d="M19.07,4.93a10,10 0 0,1 0,14.14"/>
    </svg>
  `,

  volumeMute: `
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" xmlns="http://www.w3.org/2000/svg">
      <polygon points="11,5 6,9 2,9 2,15 6,15 11,19 11,5"/>
      <line x1="15.54" y1="8.46" x2="19.07" y2="4.93"/>
      <line x1="15.54" y1="15.54" x2="19.07" y2="19.07"/>
    </svg>
  `,

  // UI 图标
  search: `
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" xmlns="http://www.w3.org/2000/svg">
      <circle cx="11" cy="11" r="8"/>
      <path d="m21,21l-4.35,-4.35"/>
    </svg>
  `,

  user: `
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" xmlns="http://www.w3.org/2000/svg">
      <path d="M20,21v-2a4,4 0 0,0 -4,-4H8a4,4 0 0,0 -4,4v2"/>
      <circle cx="12" cy="7" r="4"/>
    </svg>
  `,

  heart: `
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" xmlns="http://www.w3.org/2000/svg">
      <path d="M20.84,4.61a5.5,5.5 0 0,0 -7.78,0L12,5.67l-1.06,-1.06a5.5,5.5 0 0,0 -7.78,7.78l1.06,1.06L12,21.23l7.78,-7.78 1.06,-1.06a5.5,5.5 0 0,0 0,-7.78z"/>
    </svg>
  `,

  heartFilled: `
    <svg viewBox="0 0 24 24" fill="currentColor" xmlns="http://www.w3.org/2000/svg">
      <path d="M20.84,4.61a5.5,5.5 0 0,0 -7.78,0L12,5.67l-1.06,-1.06a5.5,5.5 0 0,0 -7.78,7.78l1.06,1.06L12,21.23l7.78,-7.78 1.06,-1.06a5.5,5.5 0 0,0 0,-7.78z"/>
    </svg>
  `,

  share: `
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" xmlns="http://www.w3.org/2000/svg">
      <path d="M18,8a3,3 0 1,0 -2.83,-4"/>
      <path d="M6,16a3,3 0 1,0 0,6a3,3 0 0,0 0,-6z"/>
      <path d="M18,16a3,3 0 1,0 0,6a3,3 0 0,0 0,-6z"/>
      <path d="M8.59,13.51l6.83,3.98"/>
      <path d="M15.41,6.51l-6.82,3.98"/>
    </svg>
  `,

  settings: `
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" xmlns="http://www.w3.org/2000/svg">
      <circle cx="12" cy="12" r="3"/>
      <path d="M19.4,15a1.65,1.65 0 0,0 .33,1.82l.06,.06a2,2 0 0,1 0,2.83a2,2 0 0,1 -2.83,0l-.06,-.06a1.65,1.65 0 0,0 -1.82,-.33a1.65,1.65 0 0,0 -1,1.51V21a2,2 0 0,1 -2,2a2,2 0 0,1 -2,-2v-.09a1.65,1.65 0 0,0 -1,-1.51a1.65,1.65 0 0,0 -1.82,.33l-.06,.06a2,2 0 0,1 -2.83,0a2,2 0 0,1 0,-2.83l.06,-.06a1.65,1.65 0 0,0 .33,-1.82a1.65,1.65 0 0,0 -1,-1.51V18a2,2 0 0,1 -2,-2a2,2 0 0,1 2,-2h.09a1.65,1.65 0 0,0 1,-1.51a1.65,1.65 0 0,0 -.33,-1.82l-.06,-.06a2,2 0 0,1 0,-2.83a2,2 0 0,1 2.83,0l.06,.06a1.65,1.65 0 0,0 1.82,.33a1.65,1.65 0 0,0 1,-1.51V3a2,2 0 0,1 2,-2a2,2 0 0,1 2,2v.09a1.65,1.65 0 0,0 1,1.51a1.65,1.65 0 0,0 1.82,-.33l.06,-.06a2,2 0 0,1 2.83,0a2,2 0 0,1 0,2.83l-.06,.06a1.65,1.65 0 0,0 -.33,1.82a1.65,1.65 0 0,0 1,1.51H21a2,2 0 0,1 2,2a2,2 0 0,1 -2,2h-.09a1.65,1.65 0 0,0 -1,1.51z"/>
    </svg>
  `,

  upload: `
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" xmlns="http://www.w3.org/2000/svg">
      <path d="M21,15v4a2,2 0 0,1 -2,2H5a2,2 0 0,1 -2,-2v-4"/>
      <polyline points="17,8 12,3 7,8"/>
      <line x1="12" y1="3" x2="12" y2="15"/>
    </svg>
  `,

  download: `
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" xmlns="http://www.w3.org/2000/svg">
      <path d="M21,15v4a2,2 0 0,1 -2,2H5a2,2 0 0,1 -2,-2v-4"/>
      <polyline points="7,16 12,21 17,16"/>
      <line x1="12" y1="3" x2="12" y2="21"/>
    </svg>
  `,

  // 主题图标
  sun: `
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" xmlns="http://www.w3.org/2000/svg">
      <circle cx="12" cy="12" r="5"/>
      <line x1="12" y1="1" x2="12" y2="3"/>
      <line x1="12" y1="21" x2="12" y2="23"/>
      <line x1="4.22" y1="4.22" x2="5.64" y2="5.64"/>
      <line x1="18.36" y1="18.36" x2="19.78" y2="19.78"/>
      <line x1="1" y1="12" x2="3" y2="12"/>
      <line x1="21" y1="12" x2="23" y2="12"/>
      <line x1="4.22" y1="19.78" x2="5.64" y2="18.36"/>
      <line x1="18.36" y1="5.64" x2="19.78" y2="4.22"/>
    </svg>
  `,

  moon: `
    <svg viewBox="0 0 24 24" fill="currentColor" xmlns="http://www.w3.org/2000/svg">
      <path d="M21,12.79A9,9 0 1,1 11.21,3a7,7 0 0,0 9.79,9.79z"/>
    </svg>
  `,

  // 导航图标
  home: `
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" xmlns="http://www.w3.org/2000/svg">
      <path d="M3,9l9,-7l9,7"/>
      <path d="M9,22V9"/>
      <path d="M15,22V9"/>
    </svg>
  `,

  library: `
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" xmlns="http://www.w3.org/2000/svg">
      <path d="M2,3h20"/>
      <path d="M6,7h16"/>
      <path d="M10,11h12"/>
      <path d="M14,15h8"/>
      <path d="M18,19h4"/>
      <path d="M2,3v18a2,2 0 0,0 2,2h4"/>
      <path d="M6,3v18a2,2 0 0,1 -2,2"/>
    </svg>
  `,

  chart: `
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" xmlns="http://www.w3.org/2000/svg">
      <line x1="18" y1="20" x2="18" y2="10"/>
      <line x1="12" y1="20" x2="12" y2="4"/>
      <line x1="6" y1="20" x2="6" y2="14"/>
    </svg>
  `,

  // 杂项图标
  close: `
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" xmlns="http://www.w3.org/2000/svg">
      <line x1="18" y1="6" x2="6" y2="18"/>
      <line x1="6" y1="6" x2="18" y2="18"/>
    </svg>
  `,

  chevronLeft: `
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" xmlns="http://www.w3.org/2000/svg">
      <polyline points="15,18 9,12 15,6"/>
    </svg>
  `,

  chevronRight: `
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" xmlns="http://www.w3.org/2000/svg">
      <polyline points="9,6 15,12 9,18"/>
    </svg>
  `,

  // 社交登录图标
  google: `
    <svg viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
      <path fill="#4285F4" d="M22.56,12.25c0,-.78 -.07,-1.53 -.2,-2.25H12v4.26h5.92c-.26,1.36 -1.04,2.51 -2.21,3.3v2.77h3.57c2.08,-1.92 3.28,-4.74 3.28,-8.08z"/>
      <path fill="#34A853" d="M12,23c2.97,0 5.46,-.98 7.28,-2.66l-3.57,-2.77c-.98,.66 -2.23,1.06 -3.71,1.06 -2.86,0 -5.29,-1.93 -6.16,-4.53H2.18v2.84C3.99,20.53 7.7,23 12,23z"/>
      <path fill="#FBBC05" d="M5.84,14.09c-.22,-.66 -.34,-1.36 -.34,-2.09s.13,-1.43 .34,-2.09V7.07H2.18C1.43,8.55 1,10.22 1,12s.43,3.45 1.18,4.93l2.85,-2.22 .81,-.62z"/>
      <path fill="#EA4335" d="M12,5.38c1.62,0 3.06,.56 4.21,1.64l3.15,-3.15C17.45,2.09 14.97,1 12,1 7.7,1 3.99,3.47 2.18,7.07l3.66,2.84c.87,-2.6 3.3,-4.53 6.16,-4.53z"/>
    </svg>
  `,

  microsoft: `
    <svg viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
      <rect x="1" y="1" width="10" height="10" fill="#F25022"/>
      <rect x="13" y="1" width="10" height="10" fill="#7FBA00"/>
      <rect x="1" y="13" width="10" height="10" fill="#00A4EF"/>
      <rect x="13" y="13" width="10" height="10" fill="#FFB900"/>
    </svg>
  `,

  apple: `
    <svg viewBox="0 0 24 24" fill="currentColor" xmlns="http://www.w3.org/2000/svg">
      <path d="M17.05,20.28c-.98,.95 -2.05,.88 -3.08,.4 -1.09,-.5 -2.08,-.48 -3.24,0 -1.44,.62 -2.2,.44 -3.06,-.4C2.79,15.25 3.51,7.59 9.05,7.31c1.35,.07 2.29,.74 3.08,.8 1.18,-.24 2.31,-.93 3.57,-.84 1.51,.12 2.65,.72 3.4,1.8 -3.12,1.87 -2.38,5.98 .48,7.13 -.57,1.5 -1.31,2.99 -2.54,4.09zm-4.31-4.08z
      <path d="M12.03,7.25c-.15,-2.23 1.66,-4.07 3.74,-4.25 .29,2.58 -2.34,4.5 -3.74,4.25z"/>
    </svg>
  `,

  wechat: `
    <svg viewBox="0 0 24 24" fill="#07C160" xmlns="http://www.w3.org/2000/svg">
      <path d="M8.691,2.188C3.891,2.188 0,5.476 0,9.53c0,2.212 1.17,4.203 3.002,5.55a.59,.59 0 0,1 .213,.665l-.39,1.48c-.019,.07 -.048,.141 -.048,.213 0,.163 .13,.295 .29,.295a.326,.326 0 0,0 .167,-.054l1.903,-1.114a.864,.864 0 0,1 .717,-.098 10.16,10.16 0 0,0 2.837,.403c.276,0 .543,-.027 .811,-.05 -.857,-2.578 .157,-4.972 1.932,-6.446 1.703,-1.415 3.882,-1.98 5.853,-1.838 -.576,-3.583 -4.196,-6.348 -8.596,-6.348z
      M5.785,5.991c.642,0 1.162,.529 1.162,1.18a1.17,1.17 0 0,1 -1.162,1.178A1.17,1.17 0 0,1 4.623,7.17c0,-.651 .52,-1.18 1.162,-1.18z
      M11.598,5.991c.642,0 1.162,.529 1.162,1.18a1.17,1.17 0 0,1 -1.162,1.178 1.17,1.17 0 0,1 -1.162,-1.178c0,-.651 .52,-1.18 1.162,-1.18z
      M15.208,14.74c-2.031,0 -3.898,.753 -5.309,1.915 -1.66,1.368 -2.562,3.275 -2.562,5.298 0,1.372 .465,2.666 1.252,3.78a.527,.527 0 0,1 .085,.396l-.229,1.098c-.032,.12 -.06,.248 -.06,.377 0,.164 .13,.296 .291,.296a.326,.326 0 0,0 .166,-.054l1.32,-.774a.774,.774 0 0,1 .642,-.087 8.44,8.44 0 0,0 2.404,.35c4.07,0 7.395,-2.91 7.395,-6.382 0,-3.472 -3.325,-6.382 -7.395,-6.382v.068z
      M13.012,16.32c.542,0 .981,.446 .981,.996a.989,.989 0 0,1 -.981,.996.989,.989 0 0,1 -.982,-.996c0,-.55 .44,-.996 .982,-.996z
      M17.404,16.32c.542,0 .982,.446 .982,.996a.989,.989 0 0,1 -.982,.996.989,.989 0 0,1 -.982,-.996c0,-.55 .44,-.996 .982,-.996z"/>
    </svg>
  `,

  qq: `
    <svg viewBox="0 0 24 24" fill="#12B7F5" xmlns="http://www.w3.org/2000/svg">
      <path d="M12,2C7.589,2 4,5.589 4,9.996c0,1.928 .691,3.691 1.835,5.07 -.18,.636 -.41,1.288 -.688,1.928 -.36,.835 -.76,1.645 -1.17,2.355 -.14,.24 -.04,.55 .21,.66 .73,.32 1.71,.56 2.63,.56 .74,0 1.43,-.14 1.98,-.39A7.94,7.94 0 0,0 12,20c1.12,0 2.18,-.23 3.15,-.64 .57,.27 1.28,.42 2.04,.42 .92,0 1.9,-.24 2.63,-.56a.488,.488 0 0,0 .21,-.66c-.41,-.71 -.81,-1.52 -1.17,-2.355 -.28,-.65 -.51,-1.31 -.69,-1.955A7.963,7.963 0 0,0 20,9.996C20,5.589 16.411,2 12,2z
      M9.5,10a1.5,1.5 0 1,1 0,-3 1.5,1.5 0 0,1 0,3z
      M14.5,10a1.5,1.5 0 1,1 0,-3 1.5,1.5 0 0,1 0,3z"/>
    </svg>
  `,
};

// Vue 图标组件工厂
export function createIcon(name, className = "") {
  return {
    name: `Icon${name.charAt(0).toUpperCase() + name.slice(1)},
    template: `<span class="${className}" v-html="Icons[name]"></span>`,
    data() {
      return { Icons };
    }
  };
}

export default Icons;
