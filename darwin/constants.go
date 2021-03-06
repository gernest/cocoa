package darwin

const (
	// Resource property flags (from CFURLPriv.h)
	// https://opensource.apple.com/source/CF/CF-635/CFURLPriv.h.auto.html
	KCFURLResourceIsRegularFile      = 0x00000001
	KCFURLResourceIsDirectory        = 0x00000002
	KCFURLResourceIsSymbolicLink     = 0x00000004
	KCFURLResourceIsVolume           = 0x00000008
	KCFURLResourceIsPackage          = 0x00000010
	KCFURLResourceIsSystemImmutable  = 0x00000020
	KCFURLResourceIsUserImmutable    = 0x00000040
	KCFURLResourceIsHidden           = 0x00000080
	KCFURLResourceHasHiddenExtension = 0x00000100
	KCFURLResourceIsApplication      = 0x00000200
	KCFURLResourceIsCompressed       = 0x00000400
	KCFURLResourceIsSystemCompressed = 0x00000400
	KCFURLCanSetHiddenExtension      = 0x00000800
	KCFURLResourceIsReadable         = 0x00001000
	KCFURLResourceIsWriteable        = 0x00002000
	KCFURLResourceIsExecutable       = 0x00004000
	KCFURLIsAliasFile                = 0x00008000
	KCFURLIsMountTrigger             = 0x00010000

	// Volume property flags (from CFURLPriv.h)
	KCFURLVolumeIsLocal               = 0x1
	KCFURLVolumeIsAutomount           = 0x2
	KCFURLVolumeDontBrowse            = 0x4
	KCFURLVolumeIsReadOnly            = 0x8
	KCFURLVolumeIsQuarantined         = 0x10
	KCFURLVolumeIsEjectable           = 0x20
	KCFURLVolumeIsRemovable           = 0x40
	KCFURLVolumeIsInternal            = 0x80
	KCFURLVolumeIsExternal            = 0x100
	KCFURLVolumeIsDiskImage           = 0x200
	KCFURLVolumeIsFileVault           = 0x400
	KCFURLVolumeIsLocaliDiskMirror    = 0x800
	KCFURLVolumeIsiPod                = 0x1000
	KCFURLVolumeIsiDisk               = 0x2000
	KCFURLVolumeIsCD                  = 0x4000
	KCFURLVolumeIsDVD                 = 0x8000
	KCFURLVolumeIsDeviceFileSystem    = 0x10000
	KCFURLVolumeSupportsPersistentIDs = 0x100000000
	KCFURLVolumeSupportsSearchFS      = 0x200000000
	KCFURLVolumeSupportsExchange      = 0x400000000
	// reserved                                           0x800000000
	KCFURLVolumeSupportsSymbolicLinks           = 0x1000000000
	KCFURLVolumeSupportsDenyModes               = 0x2000000000
	KCFURLVolumeSupportsCopyFile                = 0x4000000000
	KCFURLVolumeSupportsReadDirAttr             = 0x8000000000
	KCFURLVolumeSupportsJournaling              = 0x10000000000
	KCFURLVolumeSupportsRename                  = 0x20000000000
	KCFURLVolumeSupportsFastStatFS              = 0x40000000000
	KCFURLVolumeSupportsCaseSensitiveNames      = 0x80000000000
	KCFURLVolumeSupportsCasePreservedNames      = 0x100000000000
	KCFURLVolumeSupportsFLock                   = 0x200000000000
	KCFURLVolumeHasNoRootDirectoryTimes         = 0x400000000000
	KCFURLVolumeSupportsExtendedSecurity        = 0x800000000000
	KCFURLVolumeSupports2TBFileSize             = 0x1000000000000
	KCFURLVolumeSupportsHardLinks               = 0x2000000000000
	KCFURLVolumeSupportsMandatoryByteRangeLocks = 0x4000000000000
	KCFURLVolumeSupportsPathFromID              = 0x8000000000000
	// reserved                                      0x10000000000000
	KCFURLVolumeIsJournaling               = 0x20000000000000
	KCFURLVolumeSupportsSparseFiles        = 0x40000000000000
	KCFURLVolumeSupportsZeroRuns           = 0x80000000000000
	KCFURLVolumeSupportsVolumeSizes        = 0x100000000000000
	KCFURLVolumeSupportsRemoteEvents       = 0x200000000000000
	KCFURLVolumeSupportsHiddenFiles        = 0x400000000000000
	KCFURLVolumeSupportsDecmpFSCompression = 0x800000000000000
	KCFURLVolumeHas64BitObjectIDs          = 0x1000000000000000
	KCFURLVolumePropertyFlagsAll           = 0xffffffffffffffff

	KCFNumberSInt8Type     = 1
	KCFNumberSInt16Type    = 2
	KCFNumberSInt32Type    = 3
	KCFNumberSInt64Type    = 4
	KCFNumberFloat32Type   = 5
	KCFNumberFloat64Type   = 6
	KCFNumberCharType      = 7
	KCFNumberShortType     = 8
	KCFNumberIntType       = 9
	KCFNumberLongType      = 10
	KCFNumberLongLongType  = 11
	KCFNumberFloatType     = 12
	KCFNumberDoubleType    = 13
	KCFNumberCFIndexType   = 14
	KCFNumberNSIntegerType = 15
	KCFNumberCGFloatType   = 16
)

const (
	// FSOPT_NOFOLLOW If this bit is set, getattrlist() will not follow a
	// symlink if it occurs as the last component of path.
	FSOPT_NOFOLLOW      uint32 = 0x00000001
	FSOPT_NOINMEMUPDATE uint32 = 0x00000002
	// FSOPT_REPORT_FULLSIZE: The size of the attributes reported (in the first
	// u_int32_t field in the attribute buffer) will be the size needed to hold
	// all the requested attributes; if not set, only the attributes actu- ally
	// returned will be reported.  This allows the caller to determine if any
	// truncation occurred.
	FSOPT_REPORT_FULLSIZE uint32 = 0x00000004
	// FSOPT_PACK_INVAL_ATTRS: If this is bit is set, then all requested
	// attributes, even ones that are not supported by the object or file
	// system, will be returned. Default values will be used for the invalid
	// ones. Requires that ATTR_CMN_RETURNED_ATTRS be requested.
	FSOPT_PACK_INVAL_ATTRS uint32 = 0x00000008
	// FSOPT_ATTR_CMN_EXTENDED: If this is bit is set, then ATTR_CMN_GEN_COUNT
	// and ATTR_CMN_DOCUMENT_ID can be requested. When this option is used,
	// callers must not reference forkattrs anywhere.
	FSOPT_ATTR_CMN_EXTENDED uint32 = 0x00000020

	ATTR_CMN_NAME              uint32 = 0x00000001
	ATTR_CMN_DEVID             uint32 = 0x00000002
	ATTR_CMN_FSID              uint32 = 0x00000004
	ATTR_CMN_OBJTYPE           uint32 = 0x00000008
	ATTR_CMN_OBJTAG            uint32 = 0x00000010
	ATTR_CMN_OBJID             uint32 = 0x00000020
	ATTR_CMN_OBJPERMANENTID    uint32 = 0x00000040
	ATTR_CMN_PAROBJID          uint32 = 0x00000080
	ATTR_CMN_SCRIPT            uint32 = 0x00000100
	ATTR_CMN_CRTIME            uint32 = 0x00000200
	ATTR_CMN_MODTIME           uint32 = 0x00000400
	ATTR_CMN_CHGTIME           uint32 = 0x00000800
	ATTR_CMN_ACCTIME           uint32 = 0x00001000
	ATTR_CMN_BKUPTIME          uint32 = 0x00002000
	ATTR_CMN_FNDRINFO          uint32 = 0x00004000
	ATTR_CMN_OWNERID           uint32 = 0x00008000
	ATTR_CMN_GRPID             uint32 = 0x00010000
	ATTR_CMN_ACCESSMASK        uint32 = 0x00020000
	ATTR_CMN_FLAGS             uint32 = 0x00040000
	ATTR_CMN_USERACCESS        uint32 = 0x00200000
	ATTR_CMN_EXTENDED_SECURITY uint32 = 0x00400000
	ATTR_CMN_UUID              uint32 = 0x00800000
	ATTR_CMN_GRPUUID           uint32 = 0x01000000
	ATTR_CMN_FILEID            uint32 = 0x02000000
	ATTR_CMN_PARENTID          uint32 = 0x04000000
	ATTR_CMN_FULLPATH          uint32 = 0x08000000
	ATTR_CMN_ADDEDTIME         uint32 = 0x10000000
	ATTR_CMN_RETURNED_ATTRS    uint32 = 0x80000000
	ATTR_CMN_ALL_ATTRS         uint32 = 0x9fe7ffff

	ATTR_VOL_FSTYPE          uint32 = 0x00000001
	ATTR_VOL_SIGNATURE       uint32 = 0x00000002
	ATTR_VOL_SIZE            uint32 = 0x00000004
	ATTR_VOL_SPACEFREE       uint32 = 0x00000008
	ATTR_VOL_SPACEAVAIL      uint32 = 0x00000010
	ATTR_VOL_MINALLOCATION   uint32 = 0x00000020
	ATTR_VOL_ALLOCATIONCLUMP uint32 = 0x00000040
	ATTR_VOL_IOBLOCKSIZE     uint32 = 0x00000080
	ATTR_VOL_OBJCOUNT        uint32 = 0x00000100
	ATTR_VOL_FILECOUNT       uint32 = 0x00000200
	ATTR_VOL_DIRCOUNT        uint32 = 0x00000400
	ATTR_VOL_MAXOBJCOUNT     uint32 = 0x00000800
	ATTR_VOL_MOUNTPOINT      uint32 = 0x00001000
	ATTR_VOL_NAME            uint32 = 0x00002000
	ATTR_VOL_MOUNTFLAGS      uint32 = 0x00004000
	ATTR_VOL_MOUNTEDDEVICE   uint32 = 0x00008000
	ATTR_VOL_ENCODINGSUSED   uint32 = 0x00010000
	ATTR_VOL_CAPABILITIES    uint32 = 0x00020000
	ATTR_VOL_UUID            uint32 = 0x00040000
	ATTR_VOL_ATTRIBUTES      uint32 = 0x40000000
	ATTR_VOL_INFO            uint32 = 0x80000000
	ATTR_VOL_ALL_ATTRS       uint32 = 0xc007ffff

	ATTR_DIR_LINKCOUNT     uint32 = 0x00000001
	ATTR_DIR_ENTRYCOUNT    uint32 = 0x00000002
	ATTR_DIR_MOUNTSTATUS   uint32 = 0x00000004
	DIR_MNTSTATUS_MNTPOINT uint32 = 0x00000001
	DIR_MNTSTATUS_TRIGGER  uint32 = 0x00000002
	ATTR_DIR_ALL_ATTRS     uint32 = 0x00000007

	ATTR_FILE_LINKCOUNT     uint32 = 0x00000001
	ATTR_FILE_TOTALSIZE     uint32 = 0x00000002
	ATTR_FILE_ALLOCSIZE     uint32 = 0x00000004
	ATTR_FILE_IOBLOCKSIZE   uint32 = 0x00000008
	ATTR_FILE_DEVTYPE       uint32 = 0x00000020
	ATTR_FILE_DATALENGTH    uint32 = 0x00000200
	ATTR_FILE_DATAALLOCSIZE uint32 = 0x00000400
	ATTR_FILE_RSRCLENGTH    uint32 = 0x00001000
	ATTR_FILE_RSRCALLOCSIZE uint32 = 0x00002000

	ATTR_FILE_ALL_ATTRS uint32 = 0x0000362f

	ATTR_FORK_TOTALSIZE uint32 = 0x00000001
	ATTR_FORK_ALLOCSIZE uint32 = 0x00000002
	ATTR_FORK_ALL_ATTRS uint32 = 0x00000003

	ATTR_FILE_FORKCOUNT     uint32 = 0x00000080
	ATTR_FILE_FORKLIST      uint32 = 0x00000100
	ATTR_CMN_NAMEDATTRCOUNT uint32 = 0x00080000
	ATTR_CMN_NAMEDATTRLIST  uint32 = 0x00100000
	ATTR_FILE_DATAEXTENTS   uint32 = 0x00000800
	ATTR_FILE_RSRCEXTENTS   uint32 = 0x00004000
	ATTR_FILE_CLUMPSIZE     uint32 = 0x00000010
	ATTR_FILE_FILETYPE      uint32 = 0x00000040
)

const (
	// from sys/vnode.h
	VNON uint32 = iota
	VREG
	VDIR
	VBLK
	VCHR
	VLNK
	VSOCK
	VFIFO
	VBAD
	VSTR
	VCPLX
)

// finder flags
// https://opensource.apple.com/source/CarbonHeaders/CarbonHeaders-9A581/Finder.h
const (
	FFKIsOnDesk = 0x0001 /* Files and folders (System 6) */
	FFKColor    = 0x000E /* Files and folders */
	/* bit 0x0020 was kRequireSwitchLaunch, but is now reserved for future use*/
	FFKIsShared = 0x0040 /* Files only (Applications only) */
	/* If clear, the application needs to write to */
	/* its resource fork, and therefore cannot be */
	/* shared on a server */
	FFKHasNoINITs = 0x0080 /* Files only (Extensions/Control Panels only) */
	/* This file contains no INIT resource */
	FFKHasBeenInited = 0x0100 /* Files only */
	/* Clear if the file contains desktop database */
	/* resources ('BNDL', 'FREF', 'open', 'kind'...) */
	/* that have not been added yet. Set only by the Finder */
	/* Reserved for folders - make sure this bit is cleared for folders */
	/* bit 0x0200 was the letter bit for AOCE, but is now reserved for future use */
	FFKHasCustomIcon = 0x0400 /* Files and folders */
	FFKIsStationery  = 0x0800 /* Files only */
	FFKNameLocked    = 0x1000 /* Files and folders */
	FFKHasBundle     = 0x2000 /* Files and folders */
	/* Indicates that a file has a BNDL resource */
	/* Indicates that a folder is displayed as a package */
	FFKIsInvisible = 0x4000 /* Files and folders */
	FFKIsAlias     = 0x8000 /* Files only */
)
