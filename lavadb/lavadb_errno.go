package lavadb_client
 
 // access errno begin, range =-20000 ~ =-20999
 
 const E_ACCESS_BASE                               =-20000   // access错误base
 const E_ACCESS_MASTER_TIMEOUT                     =-20001   // 请求master超时
 const E_ACCESS_CELL_TIMEOUT                       =-20002   // 请求cell超时
 const E_ACCESS_NO_ACTIVE_CELL                     =-20003   // 没有存活的cell
 const E_ACCESS_TABLET_RO                          =-20004   // 小表只读
 const E_ACCESS_NO_DB                              =-20005   // 无该db
 const E_ACCESS_NO_TABLE                           =-20006   // 无该大表
 const E_ACCESS_NO_CELL_TABLET                     =-20007   // 小表不存在
 const E_ACCESS_NO_RANGE                           =-20008   // range范围出错
 const E_ACCESS_BIG_RECORD_SIZE                    =-20009   // value大小超过限定值
 const E_ACCESS_RECORD_ZERO_SIZE                   =-20010   // value长度为0
 const E_ACCESS_BIG_KEY_SIZE                       =-20011   // key大小超过限定值
 const E_ACCESS_KEY_ZERO_SIZE                      =-20012   // key长度为0
 const E_ACCESS_NO_TRSF_ROUTE                      =-20013   // 找不到小表串联信息
 const E_ACCESS_TABLET_CACHE_FULL                  =-20014   // 小表缓存空间满
 const E_ACCESS_INDEX_CACHE_ERROR                  =-20015   // 访问小表索引cache出错
 const E_ACCESS_CHAIN_CACHE_ERROR                  =-20016   // 访问小表串联cache出错
 const E_ACCESS_DUMP_ERROR                         =-20017   // Dump路由信息出错
 const E_ACCESS_VER_NOT_MATCH                      =-20018   // 路由版本不一致
 const E_ACCESS_REQ_PARAM_ERROR                    =-20019   // 请求参数出错
 const E_ACCESS_NEW_ASN_PKT_ERROR                  =-20020   // 分配asn packet出错
 const E_ACCESS_DCC_MQ_FULL                        =-20021   // DCC发送MQ已经满
 const E_ACCESS_CCD_MQ_FULL                        =-20022   // CCD发送MQ已经满
 const E_ACCESS_BAD_CHOICE                         =-20023   // 返回了错误的数据包
 const E_ACCESS_BAD_ECHODATA                       =-20024   // 返回了错误的echodata
 const E_ACCESS_TRSF_NO_RANGE                      =-20025   // trsf range范围出错
 const E_ACCESS_MULTI_OP_CNT_OVER_LIMIT            =-20026   // 批量请求的个数大于配置
 const E_ACCESS_MULTI_OP_SIZE_OVER_LIMIT           =-20027   // 批量请求的包大于配置
 const E_ACCESS_MULTI_CODE                         =-20028   // 批量请求编解码错误
 const E_ACCESS_MULTI_CTREAT_TASK                  =-20029   // 批量请求生成单个请求出错
 const E_ACCESS_MULTI_TIMEOUT_ONE_NORSP            =-20030   // 批量请求超时时，单个请求还没有返回
 const E_ACCESS_RECOVER_TABLET_ILL                 =-20031   // try to recover key value in tablet, but tablet is read only or freezed
 const E_ACCESS_RECOVER_READ_UNMATCH               =-20032   // recover pre read and post read is not match
 const E_ACCESS_CELLIP3_NOT_EXIST                  =-20033   // cell ip3 is null
 
 const E_ACCESS_VER_DIFF                           =-20034   // 指定key的ver和cell中此key的ver已经不一致
 const E_ACCESS_KEY_TASKLET_QUEUE_FULL             =-20035   // key的串行化队列满
 const E_ACCESS_PREFIX_MARKER_WRONG                =-20036   // list接口，前缀超过了marker
 const E_ACCESS_DELIMITER_WRONG                    =-20037   // list接口，分隔符长度大于1
 const E_ACCESS_LIST_CELL_RSP_WRONG				=-20038	 // cell在截断的情况下，未返回next_marker;
 const E_ACCESS_LIST_TOO_MUCH_SHARD				=-20039	 // cell在截断的情况下，未返回next_marker;
 
 // access errno end
 
 // cell errno begin, range =-21000 ~ 21999
 
 const E_CELL_BASE                                 =-21000   // cell错误base
 const E_CELL_BLOCK_NUMBER_ERROR                   =-21001   // 错误的block号
 const E_CELL_ALLOC_BLOCK_ERROR                    =-21002   // 无法分配block
 const E_CELL_NO_BUCKET                            =-21003   // 无对应桶
 const E_CELL_BUCKET_BUFSIZE                       =-21004   // 桶分配的bufsize太小
 const E_CELL_NO_RECORD                            =-21006   // 无对应记录
 const E_CELL_RECORD_BUFSIZE                       =-21007   // 获取记录时分配的bufsize太小
 const E_CELL_NO_TABLE                             =-21008   // 无对应的table
 const E_CELL_EXSIT_TABLE                          =-21009   // table已经存在
 const E_CELL_NO_ENOUGH_TABLE                      =-21010   // 无足够的table
 const E_CELL_DB_NOT_EXSIT                         =-21011   // db不存在
 const E_CELL_DB_EXSIT                             =-21012   // db已存在
 const E_CELL_DBID_INVALID                         =-21014   // dbid 不合法
 const E_CELL_TABLEID_INVALID                      =-21015   // tableid 不合法
 const E_CELL_DB_NOT_EMPTY                         =-21016   // db不为空
 const E_CELL_DB_BUSY                              =-21018   // 过载
 const E_CELL_WCACHE_NOT_EXSIT                     =-21019   // 写cache不存在
 const E_CELL_WCACHE_DELETED                       =-21020   // 写cache已经删除
 const E_CELL_DECODE_ERROR                         =-21030   // 解码错误
 const E_CELL_ENCODE_ERROR                         =-21031   // 解码错误
 const E_CELL_SYNCING                              =-21040   // 同步的时候，发现该记录同步中
 const E_CELL_RECORD_SIZE_ZERO                     =-21041   // value长度为0
 const E_CELL_BLOCK_NO_EXIST                       =-21050   //指定的blockno非法
 const E_CELL_BLOCK_BAD                            =-21051   //操作block失败
 const E_CELL_CHECK_SUM                            =-21052   //比对checksum失败
 const E_CELL_NO_SPACE                             =-21053   //小表writecache或者是小表容量无空间
 const E_CELL_WRITE_CACHE_SET                      =-21054   //写入小表writecache失败
 const E_CELL_REQ_PARAM                            =-21055   //请求参数非法
 const E_CELL_TABLET_SERVICE_UNOPEN                =-21056   //小表服务开关未打开，不能提供服务
 const E_CELL_NO_MEMORY                            =-21057   //内存分配失败
 const E_CELL_TABLET_WRITABLE_OK               	=-21058   //小表关闭writecache淘汰写失败
 const E_CELL_WRITE_TOKEN_GET						=-21059   //当前小表有writeobject正在进行，生成writeobject失败，拿取唯一标识失败
 const E_CELL_WRITE_TOKEN_PUT						=-21060   //释放唯一标识失败
 
 const E_CELL_WRITE_CACHE_FULL						=-21061	 //write cache full
 const E_CELL_WRITE_CACHE_LIMIT					=-21062	 //write cache apply limit
 const E_CELL_KEY_INDEX_FULL						=-21063   //big node full
 const E_CELL_INIT_STILL_ACCESS					=-21064	 //still access when init
 const E_CELL_ROUTE_ERROR							=-21065	 //路由异常
 const E_CELL_WRITE_ERROR							=-21066	 //写盘失败
 const E_CELL_READ_ERROR							=-21067	 //读盘失败
 const E_CELL_MEM_ERROR							=-21068   //内存异常
 const E_CELL_RANGE_ERROR			    			=-21069   //范围读失败
 const E_CELL_MEM_COMPACT  						=-21070
 const E_CELL_IS_DELETED  							=-21071
 const E_CELL_REF_IMMUT   							=-21074
 const E_CELL_SSTIDX_GETREF  						=-21075
 const E_CELL_NO_SHARD 							=-21076
 const E_CELL_SHARD_SERVICE_UNOPEN 				=-21077
 const E_CELL_DELETE_CACHE_SET						=-21078
 const E_CELL_SETROUTE_FAIL						=-21080
 const	E_CELL_RANGE_READ_INVALID					=-21081
 const E_CELL_CONTROL_COMPACT_ERROR				=-21082
 const E_CELL_SET_SHARD_ERROR						=-21083
 const E_CELL_SET_DATA_ERROR						=-21084
 const E_CELL_INIT_SHRAD_ERROR						=-21085
 const E_CELL_DUMP_DATA_ERROR						=-21086
 const E_CELL_GET_SPLIT_ROUTE_ERROR				=-21087
 const E_CELL_TRANSACTION_ERROR					=-21088
 const E_CELL_LEVEL_COMPACT_ERROR					=-21089
 const E_CELL_IMMU_COMPACT_ERROR					=-21090
 const E_CELL_KEY_LEN_ERROR						=-21091
 const E_CELL_VALUE_LEN_ERROR						=-21092
 const E_CELL_SETROUTE_MIN_MAX_ERROR				=-21093
 const E_CELL_DATA_NOTSORTED						=-21094
 const E_CELL_FAST_REJECT							=-21095 //过载
 const E_CELL_LIST_ALL_DEL_RECORD					=-21096
 
 // cell errno end
 
 // master errno begin, range =-22000 ~ 22999
 
 const E_MASTER_BASE                               =-22000   // 错误base
 const E_MASTER_NO_DB                              =-22001   // 无该DB
 const E_MASTER_NO_TABLE                           =-22002   // 无该TABLE
 
 const E_MASTER_NO_CELL_TABLET                     =-22003   // 无对应CELL TABLET
 const E_MASTER_TABLE_EXIST                        =-22004   // TABLE已经存在
 
 const E_MASTER_TABLET_NOT_FREE                    =-22010   // TABLET非free
 
 const E_MASTER_E_MOVED_SRC                        =-22014   // Moved, 源出错
 const E_MASTER_E_MOVED_DST                        =-22015   // Moved, 目的出错
 const E_MASTER_E_REVIVE                           =-22017   // Revive, 出错
 
 const E_MASTER_RANGE_INVALID                      =-22020   // RANGE范围出错
 
 const E_MASTER_STATUS_CHANGE                      =-22030   // TABLET状态改变
 
 const E_MASTER_ALLOC_TABLET_ERROR                 =-22040   // 无法分配TABLET
 const E_MASTER_DEL_TABLET_ERROR                   =-22041   // 无法删除TABLET
 
 const E_MASTER_CACHE_ERROR                        =-22050   // cache出错
 
 const E_MASTER_CELL_PAIR_EXISTED                  =-22060   // cell pair已经存在
 const E_MASTER_CELL_PAIR_NOT_EXIST                =-22061   // cell pair不存在
 const E_MASTER_ADD_CELL_PAIR_ERROR                =-22062   // Add cellpair异常
 
 const E_MASTER_CELL_EXPIRE                        =-22070   // cell 超时
 const E_MASTER_TABLET_EXIST_IN_PAIR               =-22080   // cell pair中已存在该tablet
 const E_MASTER_TABLET_NOT_EXIST_IN_PAIR           =-22090   // cell pair中不存在该tablet
 
 const E_MASTER_WARN_TABLET_ERROR                  =-22091   // 无效error
 
 const E_MASTER_TABLET_EXISTED                     =-22096   // 小表已经存在
 const E_MASTER_TABLET_VER_EXISTED                 =-22100   // TABLET的版本号已经存在
 const E_MASTER_TABLET_VER_NOT_EXISTED             =-22101   // TABLET的版本号不存在
 
 const E_MASTER_LOCK_SLAVE_ERROR                   =-22102   // 无效error
 const E_MASTER_LOCK_BASE                          =-22104
 const E_MASTER_VER_NOT_COINCIDE                   =-22105   // slave, access, master版本不一致or ver == 0
 const E_MASTER_LOCK_SRC_ERROR                     =-22106
 const E_MASTER_NOT_MASTER_ROLE                    =-22107   // 非master
 const E_MASTER_SLAVE_EXPIRE                       =-22108   // 请求slave超时
 const E_MASTER_MOVED_BASE                         =-22110   // 无效error
 const E_MASTER_MOVED_SLAVE                        =-22111   // 无效error 
 const E_MASTER_UNMOVE_BASE                        =-22113   // 无效error 
 const E_MASTER_UNMOVE_SLAVE                       =-22114   // 无效error 
 
 const E_MASTER_HB_ROUTE_BASE                      =-22115   // 心跳base错误
 const E_MASTER_HB_EXPIRE                          =-22116   // 心跳超时
 const E_MASTER_HB_ERROR                           =-22117   // 心跳错误
 const E_MASTER_HB_DIFF                            =-22118   // 心跳发送和接受不一致
 const E_MASTER_HB_VER_ERROR                       =-22119   // 心跳，obj版本号异常
 
 const E_MASTER_SEND_ROUTE_BASE                    =-22120   // 发送路由base错误
 const E_MASTER_SEND_ROUTE_EXPIRE                  =-22121   // 发送路由超时
 const E_MASTER_SEND_ROUTE_ERROR                   =-22122   // 发送路由错误
 const E_MASTER_SEND_ROUTE_DIFF                    =-22123   // 发送路由和接受不一致
 const E_MASTER_SEND_ROUTE_VER_ERROR               =-22124   // 发送路由，obj版本异常
 
 const E_MASTER_SEND_SLAVE_CELLPAIR_BASE           =-22126   // 想slave 发送cell pair base错误  
 const E_MASTER_SEND_SLAVE_CELLPAIR_EXPIRE         =-22127   // 向slave 发送cell pair超时
 const E_MASTER_SEND_SLAVE_ACCESSLIST_BASE         =-22128   // 向slave 发送accesslist base错误
 const E_MASTER_SEND_SLAVE_ACCESSLIST_EXPIRE       =-22129   // 向slave 发送accesslist超时
 const E_MASTER_SEND_SLAVE_DBTRSFLIST_BASE         =-22128   // 向slave 发送accesslist base错误
 const E_MASTER_SEND_SLAVE_DBTRSFLIST_EXPIRE       =-22129   // 向slave 发送accesslist超时
 
 const E_MASTER_BORN_TABLET_EXISTED                =-22130    // Born小表已经存在
 const E_MASTER_BORN_TABLET_NOT_EXISTED            =-22131    // Born小表不存在
 const E_MASTER_TABLET_NOT_BORN                    =-22132    // 小表不是Born状态
 
 const E_MASTER_INFO_ERROR                         =-22133   // master信息出错
 
 const E_MASTER_ERROR_TABLET_ERROR                 =-22134    // 置小表状态为Error异常
 const E_MASTER_REVIVE_TABLET_ERROR                =-22135    // 恢复小表异常
 
 const E_MASTER_GET_TABLET_BY_CELL_IP_ERROR        =-22136    // 通过cellip获取小表异常
 
 const E_NOT_MASTER_ROLE                           =-22137    // 角色错误 (不是Master_role)
 const E_NOT_SLAVE_ROLE                            =-22138    // 角色错误 (不是Slave_role)
 const E_CHECK_MASTER_INFO_ERROR                   =-22139    // 核对Master信息异常
 const E_SLAVE_INFO_ERROR                          =-22140    // Slave信息异常
 const E_SLAVE_VER_ERROR                           =-22141    // Slave小表路由版本异常
 
 const E_MASTER_CHECK_TRSF_SRC_ERROR               =-22142    // Trsf: check src tablet 异常
 const E_MASTER_CHECK_TRSF_DST_ERROR               =-22143    // Trsf: check dst tablet 异常
 const E_MASTER_TRSF_TYPE_ERROR                    =-22144    // Trsf: check type 异常
 const E_MASTER_TRSF_RANGE_ERROR                   =-22145    // Trsf: check range 异常
 const E_MASTER_CHECK_TRSF_ROUTE_VER_ERROR         =-22146    // Trsf: check route ver不一致
 
 const E_MASTER_TABLET_STATUS_ERROR                =-22147    // 小表状态异常
 const E_MASTER_DEAD_TABLET_ERROR                  =-22148    // 置小表状态为Dead异常
 const E_MASTER_DEAD_NUM_ERROR                     =-22149    // 小表Dead个数错误（剩下最后一份）
 const E_MASTER_REQ_PARAM_ERROR                    =-22150    // 参数出错
 const E_MASTER_TRSF_SLOTS_EXIST                   =-22151    // 迁移的串联路由已存在
 const E_MASTER_TRSF_SLOTS_NO_SPACE                =-22152    // 迁移的串联路由数超过了上限
 const E_MASTER_TRSF_VER_LENTH_ERROR               =-22153    // 校验版本号长度出错
 const E_MASTER_TRSF_TID_CID_DIFFER                =-22154    // 迁移源小表和目标小表的tid，cid不一致
 
 const E_MASTRE_CPID_ERROR                         =-22200
 const E_MASTER_CPID_EXIST                         =-22201
 const E_MASTER_CPID_STATUS_ERROR                  =-22202
 const E_MASTER_TABLET_RECYCLE_ERROR               =-22203
 const E_MASTER_REQ_TIMEOUT                        =-22204
 const E_MASTER_TABLET_NUM_ERROR                   =-22205
 const E_MASTER_SET_TABLET_ROUTE_ERROR             =-22206
 const E_MASTER_SET_TABLET_INVALID_IP              =-22207
 const E_MASTER_SET_TABLET_INVALID_STATUS          =-22208
 const E_MASTER_NEW_TABLE_INVALID_RANGE            =-22209
 const E_MASTER_NEW_TABLE_INVALID_ROUTE_STATUS     =-22210
 const E_MASTER_NEW_TABLE_ROUTE_ERROR              =-22211
 const E_MASTER_FINISH_TRSF_ROUTE_ERROR            =-22212
 const E_MASTER_GET_TABLET_BY_RANGE_ERROR          =-22213
 const E_MASTER_CLOSE_WRITE_ERROR					=-22214
 const E_MASTER_INVALID_READ_IO_SEGMENT			=-22215
 const E_MASTER_INVALID_READ_IO_INFO				=-22216
 const E_MASTER_FINISH_TRSF_ROUTE_INVALID_STATUS   =-22217
 // master errno end
 
 // dbtrsf errno begin, range =-23000 ~ 23999
 
 const E_DBTRSF_BASE                               =-23000   // dbtrsf错误base
 const E_DBTRSF_NO_TRSF_TYPE                       =-23001   // 无该迁移类型
 const E_DBTRSF_EXPIRE                             =-23002   // 超时
 const E_DBTRSF_SRC_CELL_ERROR                     =-23003   // 迁移源出错
 const E_DBTRSF_DST_CELL_ERROR                     =-23004   // 迁移目的出错
 const E_DBTRSF_MASTER_ERROR                       =-23005   // master出错
 
 const E_DBTRSF_ALREADY_DOING                      =-23007   // 迁移计划已经处于doing状态
 const E_DBTRSF_ALREADY_DDONE                      =-23008   // 迁移计划已经处于done状态
 const E_DBTRSF_ALREADY_ERROR                      =-23009   // 迁移计划已经处于error状态
 const E_DBTRSF_ALREADY_DREADY                     =-23010   // 迁移计划已经处于ready状态
 const E_DBTRSF_NO_PLAN                            =-23011   // 无效error
 const E_DBTRSF_SRC_CELL_STATUS_ERROR              =-23012   // 迁移源小表状态错误，没有可用迁移源
 
 const E_DBTRSF_MEMORY_ERROR                       =-23013   // 分配内存出错
 const E_DBTRSF_TRSF_TYPE                          =-23014   // 无效error
 const E_DBTRSF_ALLOC_SHARD_TIMEOUT               =-23015   // 分配目标小表超时
 const E_DBTRSF_NEW_TRSF_ROUTE_TIMEOUT             =-23016   // 新建串联路由超时 
 const E_DBTRSF_CHECK_TRSF_ROUTE_TIMEOUT           =-23017   // 验证master与access串联路由同步一致 超时
 const E_DBTRSF_TRSF_TIMEOUT                       =-23018   // 无效error
 const E_DBTRSF_DEL_TRSF_ROUTE_TIMEOUT      	    =-23019   // 无效error
 const E_DBTRSF_SHARD_SEND_TIMEOUT                =-23020   // 发送数据至目标小表超时
 const E_DBTRSF_SHARD_DUMP_TIMEOUT                =-23021   // 从源小表获取数据超时
 const E_DBTRSF_SHARD_DUMP_TYPE_ERROR     	    =-23022   // dump的类型出错，不是操作writecache也不是操作block
 const E_DBTRSF_CHANGE_PLAN_STATUS_ERROR   	    =-23023   // 改变trsfplan的状态出错
 const E_DBTRSF_PLAN_STATUS_ERROR                  =-23024   // 无效error
 const E_DBTRSF_PLAN_EXISTED                       =-23025   // trsfplan已经存在
 const E_DBTRSF_PLAN_CACHE_ERROR                   =-23026   // trsfplan=-cache出错
 const E_DBTRSF_BUF_CACHE_ERROR                    =-23027   // 无效error 
 const E_DBTRSF_SRC_SHARD_CHECK_TIMEOUT      	    =-23028   // 验证源小表状态 超时
 const E_DBTRSF_QUERY_PLAN     		            =-23029   // 查询迁移计划失败
 const E_DBTRSF_SEND_DST_SHARD_INDEX_ERROR   	    =-23030   // 目标小表的索引非法
 const E_DBTRSF_PLAN_NOT_EXIST   		            =-23031   // 迁移计划不存在
 const E_DBTRSF_PLAN_CACHE_GET   		            =-23032   // trsfplan=-cache出错
 const E_DBTRSF_PLAN_CACHE_FULL   		            =-23033   // trsfplan=-cache 满
 const E_DBTRSF_FINISH_TRSF_TIMEOUT      	        =-23034   // 关闭串联路由 超时
 const E_DBTRSF_SHARD_SEND_TYPE_ERROR     	    =-23035   // 发送至目标小表的数据类型出错，不是writecache的也不是block的
 const E_DBTRSF_SHARD_NO_TRSF_ROUTE   	        =-23036   // 串联路由出错无效
 const E_DBTRSF_PLAN_EXIST   		                =-23037   // 迁移计划已经存在
 const E_DBTRSF_INIT_SHARD_TIMEOUT               =-23038    // 初始化目标小表 超时
 const E_DBTRSF_CTRL_SHARD_TIMEOUT               =-23039    // 开关操作小表的writecache写淘汰 超时
 const E_DBTRSF_NO_NORMAL_SRC_IP                  =-23040    // 无效error
 const E_DBTRSF_GET_SHARD_ERROR                  =-23041    // 获取源小表错误
 const E_DBTRSF_LS_CELL_PAIR_TIMEOUT              =-23042    // 无效error
 const E_DBTRSF_LS_CELL_SHARD_TIMEOUT            =-23043    // 拉取cell小表超时
 const E_DBTRSF_ALLOC_INDEX_ERROR                 =-23044    // 分配的目标小表数组下标越界             
 const E_DBTRSF_TRSF_STEP_ERROR                   =-23045    // dbtrsf步骤跳转错误
 const E_DBTRSF_TRSF_PLAN_TYPE_ERROR              =-23046    // 找不到该迁移计划类型
 const E_DBTRSF_TRSF_PLAN_NOT_FOUND               =-23047    // 找不到该迁移任务
 const E_DBTRSF_DELETE_DOING_PLAN                 =-23048    // 用于删除正在执行任务的计划
 const E_DBTRSF_CHECK_SHARD_TIMEOUT              =-23049
 const E_DBTRSF_GET_SHARD_INFO_ERR               =-23050
 const E_DBTRSF_CHECK_SD_INFO_NE                  =-23051
 const E_DBTRSF_DUMP_SHARD_TIMEOUT               =-23052
 const E_DBTRSF_CHECK_DST_WC_KEY_ERR              =-23053
 const E_DBTRSF_GET_KV_TIMEOUT                    =-23054
 const E_DBTRSF_SEND_BLOCK_IO_ERROR               =-23055
 const E_DBTRSF_CHECK_SD_INFO_BLOCK_NE            =-23056
 const E_DBTRSF_CHECK_SD_INFO_BIG_RECORD_NE       =-23057
 const E_DBTRSF_CHECK_SD_INFO_BIG_SIZE_NE         =-23058
 const E_DBTRSF_CHECK_SD_INFO_SMALL_RECORD_NE     =-23059
 const E_DBTRSF_CHECK_SD_INFO_SMALL_SIZE_NE       =-23060
 const E_DBTRSF_GET_SPLITROUTE_TIMEOUT            =-23061
 const E_DBTRSF_GET_SPLITROUTE_ERROR              =-23062
 const E_DBTRSF_CHECK_SHARDID_NE                  =-23063
 const E_DBTRSF_CHECK_SSTBLOCK_USED_NE                  =-23064
 const E_DBTRSF_CHECK_SSTBLOCK_KV_NE                  =-23065
 const E_DBTRSF_CHECK_SSTBLOCK_DATA_TOTAL_SIZE_NE                  =-23066
 const E_DBTRSF_CHECK_SSTBLOCK_DEL_TOTAL_SIZE_NE                  =-23067
 const E_DBTRSF_CHECK_TID_NE                  =-23067
 const E_DBTRSF_CHECK_CID_NE                  =-23068
 const E_DBTRSF_ALLOC_FLAG_ERR                =-23069
 const E_DBTRSF_GET_DUMP_SHARD_ERROR          =-23070
 const E_DBTRSF_CHECK_SHARD_CONTROL_SERVER_ERROR =-23071
 const E_DBTRSF_SPLIT_BY_RANGE_PARAMS_ERROR	 	=-23072
 const E_DBTRSF_MERGE_SPLIT_PARAMS_ERROR			=-23073
 const E_DBTRSF_MERGE_ONLY_ALLOW_SRC_CPID_TABLETID		=-23074
 // dbtrsf errno end
 
 const E_PROXY_BASE          =-24000  //
 const E_PROXY_ACCESS_TIMEOUT          =-24001  //
 const E_PROXY_REQ_PARAM_ERROR                    =-24002
 const E_PROXY_NEW_ASN_PKT_ERROR                  =-24003
 const E_PROXY_DCC_MQ_FULL                        =-24004
 const E_PROXY_CCD_MQ_FULL                        =-24005
 const E_PROXY_BAD_CHOICE                         =-24006
 const E_PROXY_BAD_ECHODATA                       =-24007
 const E_PROXY_TRSF_NO_RANGE                      =-24008
 const E_PROXY_MULTI_OP_CNT_OVER_LIMIT            =-24009
 const E_PROXY_MULTI_OP_SIZE_OVER_LIMIT           =-24010
 const E_PROXY_MULTI_CODE                         =-24011
 const E_PROXY_MULTI_CTREAT_TASK                  =-24012
 const E_PROXY_MULTI_TIMEOUT_ONE_NORSP            =-24013
 const E_PROXY_BID_PARAM_ERROR                     =-24014
 const E_PROXY_CACHE_ERROR                          =-24015
 const E_PROXY_CACHE_NOT_EXIST                      =-24016
 const E_PROXY_CACHE_FULL	                        =-24017
 const E_PROXY_CACHE_NOT_PERMISSION                =-24018
 const E_PROXY_BID_NOT_EXIST   		            =-24019   //
 const E_PROXY_BID_EXIST   		            	=-24020   //
 //access_proxy errno end
 
 //semi errno begin
 const E_SEMI_BASE          						=-25000  //
 const E_SEMI_CELL_TIMEOUT          				=-25001  //
 const E_SEMI_ASN_PKT_ERROR          				=-25002  //
 const E_SEMI_BAD_CHOICE                         	=-25003
 const E_SEMI_BAD_ECHODATA                       	=-25004
 
 const E_SEMI_BUILD_REQ_ERROR                      =-25005
 const E_SEMI_TABLE_ERROR                      	=-25006
 const E_SEMI_CONDITON_EXPIRE                  	=-25007
 const E_SEMI_CONDITON_CHECK_ERR                 	=-25009
 const E_SEMI_DATA_ERR                 		=-25010
 const E_SEMI_CONDITON_TYPE_ERR                 	=-25011
 const E_SEMI_NO_TAG                 				=-25012
 const E_SEMI_NO_SUBKEY                			=-25013
 const E_SEMI_HAVE_SUBKEY                			=-25014
 
 //serialize
 const E_SEMI_SERIALIZE_RSP_ERROR                  =-25015
 const E_SEMI_HAVE_RECORD                          =-25016
 const E_SEMI_ROW_TLV_EMPTY                        =-25017
 const E_SEMI_META_DATA_RELATION                   =-25018
 const E_SEMI_META_DATA_OP_ERR                     =-25019
 
 //limit
 const E_SEMI_VALUE_LEN_ERR             			=-25020
 const E_SEMI_TLV_NUM_ERR							=-25021
 const E_SEMI_SUBKEY_LEN_ERR						=-25022
 const E_SEMI_ROW_NUM_ERR							=-25023
 const E_SEMI_SUBKEY_NUM_ERR						=-25024
 const E_SEMI_TAG_NUM_ERR							=-25025
 const E_SEMI_CONDITION_NUM_ERR					=-25026
 const E_SEMI_ORDER_NUM_ERR						=-25027
 const E_SEMI_UPDATE_NUM_ERR						=-25028
 const E_SEMI_INSERT_NUM_ERR						=-25029
 const E_SEMI_CONDITION_SET_NUM_ERR				=-25030
 const E_SEMI_TAG_REPEAT							=-25031
 const E_SEMI_SERIALIZE_PARAM_ERROR                =-25032
 
 //type
 const E_SEMI_CONDITYON_TYPE_ERR					=-25033
 const E_SEMI_DATA_TYPE_ERR						=-25034
 const E_SEMI_ORDER_DIRECTION_ERR					=-25035
 const E_SEMI_ORDER_TYPE_ERR						=-25036
 const E_SEMI_METADATA_OP_TYPE_ERR					=-25037
 const E_SEMI_METADATA_RELATION_TYPE_ERR			=-25038
 const E_SEMI_INTERNAL_TYPE_ERR					=-25039
 const E_SEMI_LIST_TYPE_ERR						=-25040
 const E_SEMI_RSP_TYPE_ERR							=-25041
 const E_SEMI_CONDITION_DEFAULT_VALUE_TYPE_ERR		=-25042
 const E_SEMI_ORDER_DEFAULT_VALUE_TYPE_ERR			=-25045
 
 // reserve
 const E_SEMI_RESERVE_ERR                          =-25043
 
 // unpress
 const E_PROXY_UNCOMPRESS_ERR                      =-25044
 
 // proxy mq full
 const E_SEMI_DCC_MQ_FULL                          =-25048
 const E_SEMI_MCD_MQ_FULL                          =-25049
 
 // ccns errno begin, range 26000 ~ 26999
 const E_CCNS_TASK_STATUS_NOT_MATCH                 =-26000
 const E_CCNS_TASK_CACHE_FULL                       =-26001
 const E_CCNS_TASK_NOT_EXIST                        =-26002
 const E_CCNS_TASK_CACHE_ERROR                      =-26003
 const E_CCNS_CP_CREATE_TMP_FILE_FAILED             =-26004
 const E_CCNS_CP_GEN_TEM_FILE_FAILED                =-26005
 const E_CCNS_CP_CFS_DIR_NOT_EXIST                  =-26006
 const E_CCNS_CP_RENAME_DST_FILE_FAILED             =-26007
 const E_CCNS_CP_COPY_TO_CFS_FAILED                 =-26008
 const E_CCNS_DUMP_DO_DUMP_FILE_FAILED              =-26009
 const E_CCNS_DUMP_OP_LOCAL_FILE_FAILED             =-26010
 const E_CCNS_DUMP_MEM_POOL_ERROR                   =-26011
 const E_CCNS_CAN_NOT_CLEAR                         =-26012
 const E_CCNS_CLEAR_CACHE_FAILED                    =-26013
 const E_CCNS_DEL_TASK_FAILED                       =-26014
 const E_CCNS_BU_RESULT_NOT_EXIST                   =-26015
 const E_CCNS_TASK_IS_EXIST                         =-26016
 const E_CCNS_CONF_PARAM_ERROR                      =-26017
 
 // lavadb cell 
 const E_CELL_GET_ERROR                             =-30002   
 const E_CELL_SET_ERROR                             =-30003   
 const E_CELL_DEL_ERROR                             =-30004   
 const E_CELL_NO_DATA                               =-30005   
 const E_CELL_SETROUTE_ERROR                        =-30006
 const E_CELL_REF_IMMUT_ERROR                       =-30007  
 const E_CELL_SST_ERROR                         =-30008      
 
 