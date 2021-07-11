// Code generated by go generate; DO NOT EDIT
// This file was generated by robots at
// 2021-07-11T10:51:09.260793
// using data from
// https://filippo.io/linux-syscall-table/TABELLA_64.json

// +build linux

package syscall

var codeToName = map[uint64]string{
  0 : "read",
  1 : "write",
  2 : "open",
  3 : "close",
  4 : "stat",
  5 : "fstat",
  6 : "lstat",
  7 : "poll",
  8 : "lseek",
  9 : "mmap",
  10 : "mprotect",
  11 : "munmap",
  12 : "brk",
  13 : "rt_sigaction",
  14 : "rt_sigprocmask",
  15 : "rt_sigreturn",
  16 : "ioctl",
  17 : "pread64",
  18 : "pwrite64",
  19 : "readv",
  20 : "writev",
  21 : "access",
  22 : "pipe",
  23 : "select",
  24 : "sched_yield",
  25 : "mremap",
  26 : "msync",
  27 : "mincore",
  28 : "madvise",
  29 : "shmget",
  30 : "shmat",
  31 : "shmctl",
  32 : "dup",
  33 : "dup2",
  34 : "pause",
  35 : "nanosleep",
  36 : "getitimer",
  37 : "alarm",
  38 : "setitimer",
  39 : "getpid",
  40 : "sendfile",
  41 : "socket",
  42 : "connect",
  43 : "accept",
  44 : "sendto",
  45 : "recvfrom",
  46 : "sendmsg",
  47 : "recvmsg",
  48 : "shutdown",
  49 : "bind",
  50 : "listen",
  51 : "getsockname",
  52 : "getpeername",
  53 : "socketpair",
  54 : "setsockopt",
  55 : "getsockopt",
  56 : "clone",
  57 : "fork",
  58 : "vfork",
  59 : "execve",
  60 : "exit",
  61 : "wait4",
  62 : "kill",
  63 : "uname",
  64 : "semget",
  65 : "semop",
  66 : "semctl",
  67 : "shmdt",
  68 : "msgget",
  69 : "msgsnd",
  70 : "msgrcv",
  71 : "msgctl",
  72 : "fcntl",
  73 : "flock",
  74 : "fsync",
  75 : "fdatasync",
  76 : "truncate",
  77 : "ftruncate",
  78 : "getdents",
  79 : "getcwd",
  80 : "chdir",
  81 : "fchdir",
  82 : "rename",
  83 : "mkdir",
  84 : "rmdir",
  85 : "creat",
  86 : "link",
  87 : "unlink",
  88 : "symlink",
  89 : "readlink",
  90 : "chmod",
  91 : "fchmod",
  92 : "chown",
  93 : "fchown",
  94 : "lchown",
  95 : "umask",
  96 : "gettimeofday",
  97 : "getrlimit",
  98 : "getrusage",
  99 : "sysinfo",
  100 : "times",
  101 : "ptrace",
  102 : "getuid",
  103 : "syslog",
  104 : "getgid",
  105 : "setuid",
  106 : "setgid",
  107 : "geteuid",
  108 : "getegid",
  109 : "setpgid",
  110 : "getppid",
  111 : "getpgrp",
  112 : "setsid",
  113 : "setreuid",
  114 : "setregid",
  115 : "getgroups",
  116 : "setgroups",
  117 : "setresuid",
  118 : "getresuid",
  119 : "setresgid",
  120 : "getresgid",
  121 : "getpgid",
  122 : "setfsuid",
  123 : "setfsgid",
  124 : "getsid",
  125 : "capget",
  126 : "capset",
  127 : "rt_sigpending",
  128 : "rt_sigtimedwait",
  129 : "rt_sigqueueinfo",
  130 : "rt_sigsuspend",
  131 : "sigaltstack",
  132 : "utime",
  133 : "mknod",
  134 : "uselib",
  135 : "personality",
  136 : "ustat",
  137 : "statfs",
  138 : "fstatfs",
  139 : "sysfs",
  140 : "getpriority",
  141 : "setpriority",
  142 : "sched_setparam",
  143 : "sched_getparam",
  144 : "sched_setscheduler",
  145 : "sched_getscheduler",
  146 : "sched_get_priority_max",
  147 : "sched_get_priority_min",
  148 : "sched_rr_get_interval",
  149 : "mlock",
  150 : "munlock",
  151 : "mlockall",
  152 : "munlockall",
  153 : "vhangup",
  154 : "modify_ldt",
  155 : "pivot_root",
  156 : "_sysctl",
  157 : "prctl",
  158 : "arch_prctl",
  159 : "adjtimex",
  160 : "setrlimit",
  161 : "chroot",
  162 : "sync",
  163 : "acct",
  164 : "settimeofday",
  165 : "mount",
  166 : "umount2",
  167 : "swapon",
  168 : "swapoff",
  169 : "reboot",
  170 : "sethostname",
  171 : "setdomainname",
  172 : "iopl",
  173 : "ioperm",
  174 : "create_module",
  175 : "init_module",
  176 : "delete_module",
  177 : "get_kernel_syms",
  178 : "query_module",
  179 : "quotactl",
  180 : "nfsservctl",
  181 : "getpmsg",
  182 : "putpmsg",
  183 : "afs_syscall",
  184 : "tuxcall",
  185 : "security",
  186 : "gettid",
  187 : "readahead",
  188 : "setxattr",
  189 : "lsetxattr",
  190 : "fsetxattr",
  191 : "getxattr",
  192 : "lgetxattr",
  193 : "fgetxattr",
  194 : "listxattr",
  195 : "llistxattr",
  196 : "flistxattr",
  197 : "removexattr",
  198 : "lremovexattr",
  199 : "fremovexattr",
  200 : "tkill",
  201 : "time",
  202 : "futex",
  203 : "sched_setaffinity",
  204 : "sched_getaffinity",
  205 : "set_thread_area",
  206 : "io_setup",
  207 : "io_destroy",
  208 : "io_getevents",
  209 : "io_submit",
  210 : "io_cancel",
  211 : "get_thread_area",
  212 : "lookup_dcookie",
  213 : "epoll_create",
  214 : "epoll_ctl_old",
  215 : "epoll_wait_old",
  216 : "remap_file_pages",
  217 : "getdents64",
  218 : "set_tid_address",
  219 : "restart_syscall",
  220 : "semtimedop",
  221 : "fadvise64",
  222 : "timer_create",
  223 : "timer_settime",
  224 : "timer_gettime",
  225 : "timer_getoverrun",
  226 : "timer_delete",
  227 : "clock_settime",
  228 : "clock_gettime",
  229 : "clock_getres",
  230 : "clock_nanosleep",
  231 : "exit_group",
  232 : "epoll_wait",
  233 : "epoll_ctl",
  234 : "tgkill",
  235 : "utimes",
  236 : "vserver",
  237 : "mbind",
  238 : "set_mempolicy",
  239 : "get_mempolicy",
  240 : "mq_open",
  241 : "mq_unlink",
  242 : "mq_timedsend",
  243 : "mq_timedreceive",
  244 : "mq_notify",
  245 : "mq_getsetattr",
  246 : "kexec_load",
  247 : "waitid",
  248 : "add_key",
  249 : "request_key",
  250 : "keyctl",
  251 : "ioprio_set",
  252 : "ioprio_get",
  253 : "inotify_init",
  254 : "inotify_add_watch",
  255 : "inotify_rm_watch",
  256 : "migrate_pages",
  257 : "openat",
  258 : "mkdirat",
  259 : "mknodat",
  260 : "fchownat",
  261 : "futimesat",
  262 : "newfstatat",
  263 : "unlinkat",
  264 : "renameat",
  265 : "linkat",
  266 : "symlinkat",
  267 : "readlinkat",
  268 : "fchmodat",
  269 : "faccessat",
  270 : "pselect6",
  271 : "ppoll",
  272 : "unshare",
  273 : "set_robust_list",
  274 : "get_robust_list",
  275 : "splice",
  276 : "tee",
  277 : "sync_file_range",
  278 : "vmsplice",
  279 : "move_pages",
  280 : "utimensat",
  281 : "epoll_pwait",
  282 : "signalfd",
  283 : "timerfd_create",
  284 : "eventfd",
  285 : "fallocate",
  286 : "timerfd_settime",
  287 : "timerfd_gettime",
  288 : "accept4",
  289 : "signalfd4",
  290 : "eventfd2",
  291 : "epoll_create1",
  292 : "dup3",
  293 : "pipe2",
  294 : "inotify_init1",
  295 : "preadv",
  296 : "pwritev",
  297 : "rt_tgsigqueueinfo",
  298 : "perf_event_open",
  299 : "recvmmsg",
  300 : "fanotify_init",
  301 : "fanotify_mark",
  302 : "prlimit64",
  303 : "name_to_handle_at",
  304 : "open_by_handle_at",
  305 : "clock_adjtime",
  306 : "syncfs",
  307 : "sendmmsg",
  308 : "setns",
  309 : "getcpu",
  310 : "process_vm_readv",
  311 : "process_vm_writev",
  312 : "kcmp",
  313 : "finit_module",
}
var nameToCode = map[string]uint64{
  "read" : 0,
  "write" : 1,
  "open" : 2,
  "close" : 3,
  "stat" : 4,
  "fstat" : 5,
  "lstat" : 6,
  "poll" : 7,
  "lseek" : 8,
  "mmap" : 9,
  "mprotect" : 10,
  "munmap" : 11,
  "brk" : 12,
  "rt_sigaction" : 13,
  "rt_sigprocmask" : 14,
  "rt_sigreturn" : 15,
  "ioctl" : 16,
  "pread64" : 17,
  "pwrite64" : 18,
  "readv" : 19,
  "writev" : 20,
  "access" : 21,
  "pipe" : 22,
  "select" : 23,
  "sched_yield" : 24,
  "mremap" : 25,
  "msync" : 26,
  "mincore" : 27,
  "madvise" : 28,
  "shmget" : 29,
  "shmat" : 30,
  "shmctl" : 31,
  "dup" : 32,
  "dup2" : 33,
  "pause" : 34,
  "nanosleep" : 35,
  "getitimer" : 36,
  "alarm" : 37,
  "setitimer" : 38,
  "getpid" : 39,
  "sendfile" : 40,
  "socket" : 41,
  "connect" : 42,
  "accept" : 43,
  "sendto" : 44,
  "recvfrom" : 45,
  "sendmsg" : 46,
  "recvmsg" : 47,
  "shutdown" : 48,
  "bind" : 49,
  "listen" : 50,
  "getsockname" : 51,
  "getpeername" : 52,
  "socketpair" : 53,
  "setsockopt" : 54,
  "getsockopt" : 55,
  "clone" : 56,
  "fork" : 57,
  "vfork" : 58,
  "execve" : 59,
  "exit" : 60,
  "wait4" : 61,
  "kill" : 62,
  "uname" : 63,
  "semget" : 64,
  "semop" : 65,
  "semctl" : 66,
  "shmdt" : 67,
  "msgget" : 68,
  "msgsnd" : 69,
  "msgrcv" : 70,
  "msgctl" : 71,
  "fcntl" : 72,
  "flock" : 73,
  "fsync" : 74,
  "fdatasync" : 75,
  "truncate" : 76,
  "ftruncate" : 77,
  "getdents" : 78,
  "getcwd" : 79,
  "chdir" : 80,
  "fchdir" : 81,
  "rename" : 82,
  "mkdir" : 83,
  "rmdir" : 84,
  "creat" : 85,
  "link" : 86,
  "unlink" : 87,
  "symlink" : 88,
  "readlink" : 89,
  "chmod" : 90,
  "fchmod" : 91,
  "chown" : 92,
  "fchown" : 93,
  "lchown" : 94,
  "umask" : 95,
  "gettimeofday" : 96,
  "getrlimit" : 97,
  "getrusage" : 98,
  "sysinfo" : 99,
  "times" : 100,
  "ptrace" : 101,
  "getuid" : 102,
  "syslog" : 103,
  "getgid" : 104,
  "setuid" : 105,
  "setgid" : 106,
  "geteuid" : 107,
  "getegid" : 108,
  "setpgid" : 109,
  "getppid" : 110,
  "getpgrp" : 111,
  "setsid" : 112,
  "setreuid" : 113,
  "setregid" : 114,
  "getgroups" : 115,
  "setgroups" : 116,
  "setresuid" : 117,
  "getresuid" : 118,
  "setresgid" : 119,
  "getresgid" : 120,
  "getpgid" : 121,
  "setfsuid" : 122,
  "setfsgid" : 123,
  "getsid" : 124,
  "capget" : 125,
  "capset" : 126,
  "rt_sigpending" : 127,
  "rt_sigtimedwait" : 128,
  "rt_sigqueueinfo" : 129,
  "rt_sigsuspend" : 130,
  "sigaltstack" : 131,
  "utime" : 132,
  "mknod" : 133,
  "uselib" : 134,
  "personality" : 135,
  "ustat" : 136,
  "statfs" : 137,
  "fstatfs" : 138,
  "sysfs" : 139,
  "getpriority" : 140,
  "setpriority" : 141,
  "sched_setparam" : 142,
  "sched_getparam" : 143,
  "sched_setscheduler" : 144,
  "sched_getscheduler" : 145,
  "sched_get_priority_max" : 146,
  "sched_get_priority_min" : 147,
  "sched_rr_get_interval" : 148,
  "mlock" : 149,
  "munlock" : 150,
  "mlockall" : 151,
  "munlockall" : 152,
  "vhangup" : 153,
  "modify_ldt" : 154,
  "pivot_root" : 155,
  "_sysctl" : 156,
  "prctl" : 157,
  "arch_prctl" : 158,
  "adjtimex" : 159,
  "setrlimit" : 160,
  "chroot" : 161,
  "sync" : 162,
  "acct" : 163,
  "settimeofday" : 164,
  "mount" : 165,
  "umount2" : 166,
  "swapon" : 167,
  "swapoff" : 168,
  "reboot" : 169,
  "sethostname" : 170,
  "setdomainname" : 171,
  "iopl" : 172,
  "ioperm" : 173,
  "create_module" : 174,
  "init_module" : 175,
  "delete_module" : 176,
  "get_kernel_syms" : 177,
  "query_module" : 178,
  "quotactl" : 179,
  "nfsservctl" : 180,
  "getpmsg" : 181,
  "putpmsg" : 182,
  "afs_syscall" : 183,
  "tuxcall" : 184,
  "security" : 185,
  "gettid" : 186,
  "readahead" : 187,
  "setxattr" : 188,
  "lsetxattr" : 189,
  "fsetxattr" : 190,
  "getxattr" : 191,
  "lgetxattr" : 192,
  "fgetxattr" : 193,
  "listxattr" : 194,
  "llistxattr" : 195,
  "flistxattr" : 196,
  "removexattr" : 197,
  "lremovexattr" : 198,
  "fremovexattr" : 199,
  "tkill" : 200,
  "time" : 201,
  "futex" : 202,
  "sched_setaffinity" : 203,
  "sched_getaffinity" : 204,
  "set_thread_area" : 205,
  "io_setup" : 206,
  "io_destroy" : 207,
  "io_getevents" : 208,
  "io_submit" : 209,
  "io_cancel" : 210,
  "get_thread_area" : 211,
  "lookup_dcookie" : 212,
  "epoll_create" : 213,
  "epoll_ctl_old" : 214,
  "epoll_wait_old" : 215,
  "remap_file_pages" : 216,
  "getdents64" : 217,
  "set_tid_address" : 218,
  "restart_syscall" : 219,
  "semtimedop" : 220,
  "fadvise64" : 221,
  "timer_create" : 222,
  "timer_settime" : 223,
  "timer_gettime" : 224,
  "timer_getoverrun" : 225,
  "timer_delete" : 226,
  "clock_settime" : 227,
  "clock_gettime" : 228,
  "clock_getres" : 229,
  "clock_nanosleep" : 230,
  "exit_group" : 231,
  "epoll_wait" : 232,
  "epoll_ctl" : 233,
  "tgkill" : 234,
  "utimes" : 235,
  "vserver" : 236,
  "mbind" : 237,
  "set_mempolicy" : 238,
  "get_mempolicy" : 239,
  "mq_open" : 240,
  "mq_unlink" : 241,
  "mq_timedsend" : 242,
  "mq_timedreceive" : 243,
  "mq_notify" : 244,
  "mq_getsetattr" : 245,
  "kexec_load" : 246,
  "waitid" : 247,
  "add_key" : 248,
  "request_key" : 249,
  "keyctl" : 250,
  "ioprio_set" : 251,
  "ioprio_get" : 252,
  "inotify_init" : 253,
  "inotify_add_watch" : 254,
  "inotify_rm_watch" : 255,
  "migrate_pages" : 256,
  "openat" : 257,
  "mkdirat" : 258,
  "mknodat" : 259,
  "fchownat" : 260,
  "futimesat" : 261,
  "newfstatat" : 262,
  "unlinkat" : 263,
  "renameat" : 264,
  "linkat" : 265,
  "symlinkat" : 266,
  "readlinkat" : 267,
  "fchmodat" : 268,
  "faccessat" : 269,
  "pselect6" : 270,
  "ppoll" : 271,
  "unshare" : 272,
  "set_robust_list" : 273,
  "get_robust_list" : 274,
  "splice" : 275,
  "tee" : 276,
  "sync_file_range" : 277,
  "vmsplice" : 278,
  "move_pages" : 279,
  "utimensat" : 280,
  "epoll_pwait" : 281,
  "signalfd" : 282,
  "timerfd_create" : 283,
  "eventfd" : 284,
  "fallocate" : 285,
  "timerfd_settime" : 286,
  "timerfd_gettime" : 287,
  "accept4" : 288,
  "signalfd4" : 289,
  "eventfd2" : 290,
  "epoll_create1" : 291,
  "dup3" : 292,
  "pipe2" : 293,
  "inotify_init1" : 294,
  "preadv" : 295,
  "pwritev" : 296,
  "rt_tgsigqueueinfo" : 297,
  "perf_event_open" : 298,
  "recvmmsg" : 299,
  "fanotify_init" : 300,
  "fanotify_mark" : 301,
  "prlimit64" : 302,
  "name_to_handle_at" : 303,
  "open_by_handle_at" : 304,
  "clock_adjtime" : 305,
  "syncfs" : 306,
  "sendmmsg" : 307,
  "setns" : 308,
  "getcpu" : 309,
  "process_vm_readv" : 310,
  "process_vm_writev" : 311,
  "kcmp" : 312,
  "finit_module" : 313,
}

func GetSyscallName(code uint64) string {
  return codeToName[code]
}

