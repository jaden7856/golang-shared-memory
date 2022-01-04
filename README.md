## hared Memory(공유 메모리)

PIPE, Named PIPE, Message Queue가 통신을 이용한 설비라면, **Shared Memory**는 **공유메모리가 데이터 자체를 공유하도록** **지원**하는 설비입니다.

프로세스는 자신만의 메모리 영역을 가지고 있습니다. 이 메모리 영역은 다른 프로세스가 접근해서 함부로 데이터를 읽거나 쓰지 못하도록 커널에 의해서 보호가 되는데, 만약 다른 다른 프로세스의 메모리 영역을 침범하려고 하면 커널은 침범 프로세스에 SIGSEGV(경고 시그널 - 할당된 메모리의 범위를 벗어나는곳에서 읽거나, 쓰기를 시도할 때 발생) 을 보내게 됩니다.

- **프로세스간 메모리 영역을 공유해서 사용할 수 있도록** **허용**
- 공유메모리는 중개자가 없이 곧바로 메모리에 접근할 수 있기 때문에 다른 모든 **IPC**들 중에서 가장 빠르게 **작동**할 수 있습니다.
- 공유변수에 접근할 수 있는 Go 루틴의 최대 수는 1개.
  - 두개 이상의 Go루틴이 이 변수를 동시에 업데이트할 수 없다



### 메모리 제한: Windows 10

다음 표에서는 Windows 10 실제 메모리에 대한 제한을 지정합니다.

| 버전                            | X86 제한 | X64 제한 |
| ------------------------------- | -------- | -------- |
| Windows 10 Enterprise           | 4GB      | 6 TB     |
| Windows 10 Education            | 4GB      | 2TB      |
| Windows 10 Pro for Workstations | 4GB      | 6 TB     |
| Windows 10 Pro                  | 4GB      | 2TB      |
| Windows 10 Home                 | 4GB      | 128GB    |



### 메모리 제한: Linux

다음 텍스트는 **ipcs** 명령 출력의 예로, `//` 이후에 매개변수 이름을 표시하는 주석이 추가되어 있습니다.

```
# ipcs -l

------ Messages Limits --------
max queues system wide = 32000
max size of message (bytes) = 8192
default max size of queue (bytes) = 16384

------ Shared Memory Limits --------
max number of segments = 4096							# SHMMNI
max seg size (kbytes) = 18014398509465599				# SHMMAX
max total shared memory (kbytes) = 18014398442373116	# SHMALL
min seg size (bytes) = 1

------ Semaphore Limits --------
max number of arrays = 128								# SEMMNI
max semaphores per array = 250							# SEMMSL
max semaphores system wide = 32000						# SEMMNS
max ops per semop call = 32								# SEMOPM
semaphore max value = 32767
```



- 공유 메모리 한계의 첫 번째 섹션에서 시작하여 **SHMMAX**한계는  Linux 시스템에 있는 공유 메모리 세그먼트의 최대 크기입니다. **SHMALL** 한계는 시스템에 있는 공유 메모리 페이지의 최대 할당입니다.
  - **SHMMAX** 값을 시스템의 실제 메모리 양과 동일하게 설정하는 것이 좋습니다. 하지만 x86 시스템에 필요한 최소량은 268435456(256MB)이고 64비트 시스템의 경우에는 1073741824(1GB)입니다.
- 다음 섹션에서는 운영 체제에 사용 가능한 세마포어 양에 대해 다룹니다. 커널 매개변수 **sem**은 네 개의 토큰(**SEMMSL**, **SEMMNS**, **SEMOPM** 및 **SEMMNI**)으로 구성됩니다. **SEMMNS**는 **SEMMSL**에 **SEMMNI**를 곱한 결과입니다. 데이터베이스 관리자에서는 필요에 따라 배열 수(**SEMMNI**)를 늘려야 합니다. 일반적으로, **SEMMNI**는 데이터 서버 컴퓨터의 논리적 파티션 수를 곱하고 데이터베이스 서버 컴퓨터의 로컬 애플리케이션 연결 수를 더한 시스템에서 예상되는 최대 에이전트 수의 두 배여야 합니다.
- 세 번째 섹션에서는 시스템에 대한 메시지에 대해 다룹니다.
  - **MSGMNI** 매개변수는 시작할 수 있는 에이전트 수에 영향을 미칩니다. **MSGMAX** 매개변수는 큐에서 전송할 수 있는 메시지의 크기에 영향을 미치며, **MSGMNB** 매개변수는 큐의 크기에 영향을 미칩니다.
  - **MSGMAX** 매개변수를 64KB(즉, 65536바이트)로 변경하고, **MSGMNB** 매개변수를 65536으로 늘려야 합니다.

## Memory Map

 Memory Map도 Shared Memory(공유메모리)공간과 마찬가지로 메모리를 공유한다는 측면에 있어서는 서로 비슷한 측면이 있습니다. 차이점은 Memory Map의 경우 **열린파일을 메모리에 맵핑시켜서,** **공유**한다는 점

