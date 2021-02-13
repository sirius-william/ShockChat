// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: Message.proto

#ifndef GOOGLE_PROTOBUF_INCLUDED_Message_2eproto
#define GOOGLE_PROTOBUF_INCLUDED_Message_2eproto

#include <limits>
#include <string>

#include <google/protobuf/port_def.inc>
#if PROTOBUF_VERSION < 3014000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers. Please update
#error your headers.
#endif
#if 3014000 < PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers. Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/port_undef.inc>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_table_driven.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/metadata_lite.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/unknown_field_set.h>
#include <google/protobuf/timestamp.pb.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
#define PROTOBUF_INTERNAL_EXPORT_Message_2eproto
PROTOBUF_NAMESPACE_OPEN
namespace internal {
class AnyMetadata;
}  // namespace internal
PROTOBUF_NAMESPACE_CLOSE

// Internal implementation detail -- do not use these members.
struct TableStruct_Message_2eproto {
  static const ::PROTOBUF_NAMESPACE_ID::internal::ParseTableField entries[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::PROTOBUF_NAMESPACE_ID::internal::AuxiliaryParseTableField aux[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::PROTOBUF_NAMESPACE_ID::internal::ParseTable schema[2]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::PROTOBUF_NAMESPACE_ID::internal::FieldMetadata field_metadata[];
  static const ::PROTOBUF_NAMESPACE_ID::internal::SerializationTable serialization_table[];
  static const ::PROTOBUF_NAMESPACE_ID::uint32 offsets[];
};
extern const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable descriptor_table_Message_2eproto;
namespace protos {
class Message;
class MessageDefaultTypeInternal;
extern MessageDefaultTypeInternal _Message_default_instance_;
class Messages;
class MessagesDefaultTypeInternal;
extern MessagesDefaultTypeInternal _Messages_default_instance_;
}  // namespace protos
PROTOBUF_NAMESPACE_OPEN
template<> ::protos::Message* Arena::CreateMaybeMessage<::protos::Message>(Arena*);
template<> ::protos::Messages* Arena::CreateMaybeMessage<::protos::Messages>(Arena*);
PROTOBUF_NAMESPACE_CLOSE
namespace protos {

// ===================================================================

class Message PROTOBUF_FINAL :
    public ::PROTOBUF_NAMESPACE_ID::Message /* @@protoc_insertion_point(class_definition:protos.Message) */ {
 public:
  inline Message() : Message(nullptr) {}
  virtual ~Message();

  Message(const Message& from);
  Message(Message&& from) noexcept
    : Message() {
    *this = ::std::move(from);
  }

  inline Message& operator=(const Message& from) {
    CopyFrom(from);
    return *this;
  }
  inline Message& operator=(Message&& from) noexcept {
    if (GetArena() == from.GetArena()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }

  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* descriptor() {
    return GetDescriptor();
  }
  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* GetDescriptor() {
    return GetMetadataStatic().descriptor;
  }
  static const ::PROTOBUF_NAMESPACE_ID::Reflection* GetReflection() {
    return GetMetadataStatic().reflection;
  }
  static const Message& default_instance();

  static inline const Message* internal_default_instance() {
    return reinterpret_cast<const Message*>(
               &_Message_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  friend void swap(Message& a, Message& b) {
    a.Swap(&b);
  }
  inline void Swap(Message* other) {
    if (other == this) return;
    if (GetArena() == other->GetArena()) {
      InternalSwap(other);
    } else {
      ::PROTOBUF_NAMESPACE_ID::internal::GenericSwap(this, other);
    }
  }
  void UnsafeArenaSwap(Message* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetArena() == other->GetArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  inline Message* New() const final {
    return CreateMaybeMessage<Message>(nullptr);
  }

  Message* New(::PROTOBUF_NAMESPACE_ID::Arena* arena) const final {
    return CreateMaybeMessage<Message>(arena);
  }
  void CopyFrom(const ::PROTOBUF_NAMESPACE_ID::Message& from) final;
  void MergeFrom(const ::PROTOBUF_NAMESPACE_ID::Message& from) final;
  void CopyFrom(const Message& from);
  void MergeFrom(const Message& from);
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  const char* _InternalParse(const char* ptr, ::PROTOBUF_NAMESPACE_ID::internal::ParseContext* ctx) final;
  ::PROTOBUF_NAMESPACE_ID::uint8* _InternalSerialize(
      ::PROTOBUF_NAMESPACE_ID::uint8* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const final;
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  inline void SharedCtor();
  inline void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(Message* other);
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "protos.Message";
  }
  protected:
  explicit Message(::PROTOBUF_NAMESPACE_ID::Arena* arena);
  private:
  static void ArenaDtor(void* object);
  inline void RegisterArenaDtor(::PROTOBUF_NAMESPACE_ID::Arena* arena);
  public:

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;
  private:
  static ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadataStatic() {
    ::PROTOBUF_NAMESPACE_ID::internal::AssignDescriptors(&::descriptor_table_Message_2eproto);
    return ::descriptor_table_Message_2eproto.file_level_metadata[kIndexInFileMessages];
  }

  public:

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kMsgFieldNumber = 3,
    kSendTimeFieldNumber = 2,
    kUseridFieldNumber = 1,
  };
  // bytes msg = 3;
  void clear_msg();
  const std::string& msg() const;
  void set_msg(const std::string& value);
  void set_msg(std::string&& value);
  void set_msg(const char* value);
  void set_msg(const void* value, size_t size);
  std::string* mutable_msg();
  std::string* release_msg();
  void set_allocated_msg(std::string* msg);
  private:
  const std::string& _internal_msg() const;
  void _internal_set_msg(const std::string& value);
  std::string* _internal_mutable_msg();
  public:

  // .google.protobuf.Timestamp sendTime = 2;
  bool has_sendtime() const;
  private:
  bool _internal_has_sendtime() const;
  public:
  void clear_sendtime();
  const PROTOBUF_NAMESPACE_ID::Timestamp& sendtime() const;
  PROTOBUF_NAMESPACE_ID::Timestamp* release_sendtime();
  PROTOBUF_NAMESPACE_ID::Timestamp* mutable_sendtime();
  void set_allocated_sendtime(PROTOBUF_NAMESPACE_ID::Timestamp* sendtime);
  private:
  const PROTOBUF_NAMESPACE_ID::Timestamp& _internal_sendtime() const;
  PROTOBUF_NAMESPACE_ID::Timestamp* _internal_mutable_sendtime();
  public:
  void unsafe_arena_set_allocated_sendtime(
      PROTOBUF_NAMESPACE_ID::Timestamp* sendtime);
  PROTOBUF_NAMESPACE_ID::Timestamp* unsafe_arena_release_sendtime();

  // int32 userid = 1;
  void clear_userid();
  ::PROTOBUF_NAMESPACE_ID::int32 userid() const;
  void set_userid(::PROTOBUF_NAMESPACE_ID::int32 value);
  private:
  ::PROTOBUF_NAMESPACE_ID::int32 _internal_userid() const;
  void _internal_set_userid(::PROTOBUF_NAMESPACE_ID::int32 value);
  public:

  // @@protoc_insertion_point(class_scope:protos.Message)
 private:
  class _Internal;

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  ::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr msg_;
  PROTOBUF_NAMESPACE_ID::Timestamp* sendtime_;
  ::PROTOBUF_NAMESPACE_ID::int32 userid_;
  mutable ::PROTOBUF_NAMESPACE_ID::internal::CachedSize _cached_size_;
  friend struct ::TableStruct_Message_2eproto;
};
// -------------------------------------------------------------------

class Messages PROTOBUF_FINAL :
    public ::PROTOBUF_NAMESPACE_ID::Message /* @@protoc_insertion_point(class_definition:protos.Messages) */ {
 public:
  inline Messages() : Messages(nullptr) {}
  virtual ~Messages();

  Messages(const Messages& from);
  Messages(Messages&& from) noexcept
    : Messages() {
    *this = ::std::move(from);
  }

  inline Messages& operator=(const Messages& from) {
    CopyFrom(from);
    return *this;
  }
  inline Messages& operator=(Messages&& from) noexcept {
    if (GetArena() == from.GetArena()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }

  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* descriptor() {
    return GetDescriptor();
  }
  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* GetDescriptor() {
    return GetMetadataStatic().descriptor;
  }
  static const ::PROTOBUF_NAMESPACE_ID::Reflection* GetReflection() {
    return GetMetadataStatic().reflection;
  }
  static const Messages& default_instance();

  static inline const Messages* internal_default_instance() {
    return reinterpret_cast<const Messages*>(
               &_Messages_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    1;

  friend void swap(Messages& a, Messages& b) {
    a.Swap(&b);
  }
  inline void Swap(Messages* other) {
    if (other == this) return;
    if (GetArena() == other->GetArena()) {
      InternalSwap(other);
    } else {
      ::PROTOBUF_NAMESPACE_ID::internal::GenericSwap(this, other);
    }
  }
  void UnsafeArenaSwap(Messages* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetArena() == other->GetArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  inline Messages* New() const final {
    return CreateMaybeMessage<Messages>(nullptr);
  }

  Messages* New(::PROTOBUF_NAMESPACE_ID::Arena* arena) const final {
    return CreateMaybeMessage<Messages>(arena);
  }
  void CopyFrom(const ::PROTOBUF_NAMESPACE_ID::Message& from) final;
  void MergeFrom(const ::PROTOBUF_NAMESPACE_ID::Message& from) final;
  void CopyFrom(const Messages& from);
  void MergeFrom(const Messages& from);
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  const char* _InternalParse(const char* ptr, ::PROTOBUF_NAMESPACE_ID::internal::ParseContext* ctx) final;
  ::PROTOBUF_NAMESPACE_ID::uint8* _InternalSerialize(
      ::PROTOBUF_NAMESPACE_ID::uint8* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const final;
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  inline void SharedCtor();
  inline void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(Messages* other);
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "protos.Messages";
  }
  protected:
  explicit Messages(::PROTOBUF_NAMESPACE_ID::Arena* arena);
  private:
  static void ArenaDtor(void* object);
  inline void RegisterArenaDtor(::PROTOBUF_NAMESPACE_ID::Arena* arena);
  public:

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;
  private:
  static ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadataStatic() {
    ::PROTOBUF_NAMESPACE_ID::internal::AssignDescriptors(&::descriptor_table_Message_2eproto);
    return ::descriptor_table_Message_2eproto.file_level_metadata[kIndexInFileMessages];
  }

  public:

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kMsgFieldNumber = 2,
    kStatusFieldNumber = 1,
  };
  // repeated .protos.Message msg = 2;
  int msg_size() const;
  private:
  int _internal_msg_size() const;
  public:
  void clear_msg();
  ::protos::Message* mutable_msg(int index);
  ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::protos::Message >*
      mutable_msg();
  private:
  const ::protos::Message& _internal_msg(int index) const;
  ::protos::Message* _internal_add_msg();
  public:
  const ::protos::Message& msg(int index) const;
  ::protos::Message* add_msg();
  const ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::protos::Message >&
      msg() const;

  // int32 status = 1;
  void clear_status();
  ::PROTOBUF_NAMESPACE_ID::int32 status() const;
  void set_status(::PROTOBUF_NAMESPACE_ID::int32 value);
  private:
  ::PROTOBUF_NAMESPACE_ID::int32 _internal_status() const;
  void _internal_set_status(::PROTOBUF_NAMESPACE_ID::int32 value);
  public:

  // @@protoc_insertion_point(class_scope:protos.Messages)
 private:
  class _Internal;

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::protos::Message > msg_;
  ::PROTOBUF_NAMESPACE_ID::int32 status_;
  mutable ::PROTOBUF_NAMESPACE_ID::internal::CachedSize _cached_size_;
  friend struct ::TableStruct_Message_2eproto;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// Message

// int32 userid = 1;
inline void Message::clear_userid() {
  userid_ = 0;
}
inline ::PROTOBUF_NAMESPACE_ID::int32 Message::_internal_userid() const {
  return userid_;
}
inline ::PROTOBUF_NAMESPACE_ID::int32 Message::userid() const {
  // @@protoc_insertion_point(field_get:protos.Message.userid)
  return _internal_userid();
}
inline void Message::_internal_set_userid(::PROTOBUF_NAMESPACE_ID::int32 value) {
  
  userid_ = value;
}
inline void Message::set_userid(::PROTOBUF_NAMESPACE_ID::int32 value) {
  _internal_set_userid(value);
  // @@protoc_insertion_point(field_set:protos.Message.userid)
}

// .google.protobuf.Timestamp sendTime = 2;
inline bool Message::_internal_has_sendtime() const {
  return this != internal_default_instance() && sendtime_ != nullptr;
}
inline bool Message::has_sendtime() const {
  return _internal_has_sendtime();
}
inline const PROTOBUF_NAMESPACE_ID::Timestamp& Message::_internal_sendtime() const {
  const PROTOBUF_NAMESPACE_ID::Timestamp* p = sendtime_;
  return p != nullptr ? *p : reinterpret_cast<const PROTOBUF_NAMESPACE_ID::Timestamp&>(
      PROTOBUF_NAMESPACE_ID::_Timestamp_default_instance_);
}
inline const PROTOBUF_NAMESPACE_ID::Timestamp& Message::sendtime() const {
  // @@protoc_insertion_point(field_get:protos.Message.sendTime)
  return _internal_sendtime();
}
inline void Message::unsafe_arena_set_allocated_sendtime(
    PROTOBUF_NAMESPACE_ID::Timestamp* sendtime) {
  if (GetArena() == nullptr) {
    delete reinterpret_cast<::PROTOBUF_NAMESPACE_ID::MessageLite*>(sendtime_);
  }
  sendtime_ = sendtime;
  if (sendtime) {
    
  } else {
    
  }
  // @@protoc_insertion_point(field_unsafe_arena_set_allocated:protos.Message.sendTime)
}
inline PROTOBUF_NAMESPACE_ID::Timestamp* Message::release_sendtime() {
  
  PROTOBUF_NAMESPACE_ID::Timestamp* temp = sendtime_;
  sendtime_ = nullptr;
  if (GetArena() != nullptr) {
    temp = ::PROTOBUF_NAMESPACE_ID::internal::DuplicateIfNonNull(temp);
  }
  return temp;
}
inline PROTOBUF_NAMESPACE_ID::Timestamp* Message::unsafe_arena_release_sendtime() {
  // @@protoc_insertion_point(field_release:protos.Message.sendTime)
  
  PROTOBUF_NAMESPACE_ID::Timestamp* temp = sendtime_;
  sendtime_ = nullptr;
  return temp;
}
inline PROTOBUF_NAMESPACE_ID::Timestamp* Message::_internal_mutable_sendtime() {
  
  if (sendtime_ == nullptr) {
    auto* p = CreateMaybeMessage<PROTOBUF_NAMESPACE_ID::Timestamp>(GetArena());
    sendtime_ = p;
  }
  return sendtime_;
}
inline PROTOBUF_NAMESPACE_ID::Timestamp* Message::mutable_sendtime() {
  // @@protoc_insertion_point(field_mutable:protos.Message.sendTime)
  return _internal_mutable_sendtime();
}
inline void Message::set_allocated_sendtime(PROTOBUF_NAMESPACE_ID::Timestamp* sendtime) {
  ::PROTOBUF_NAMESPACE_ID::Arena* message_arena = GetArena();
  if (message_arena == nullptr) {
    delete reinterpret_cast< ::PROTOBUF_NAMESPACE_ID::MessageLite*>(sendtime_);
  }
  if (sendtime) {
    ::PROTOBUF_NAMESPACE_ID::Arena* submessage_arena =
      reinterpret_cast<::PROTOBUF_NAMESPACE_ID::MessageLite*>(sendtime)->GetArena();
    if (message_arena != submessage_arena) {
      sendtime = ::PROTOBUF_NAMESPACE_ID::internal::GetOwnedMessage(
          message_arena, sendtime, submessage_arena);
    }
    
  } else {
    
  }
  sendtime_ = sendtime;
  // @@protoc_insertion_point(field_set_allocated:protos.Message.sendTime)
}

// bytes msg = 3;
inline void Message::clear_msg() {
  msg_.ClearToEmpty();
}
inline const std::string& Message::msg() const {
  // @@protoc_insertion_point(field_get:protos.Message.msg)
  return _internal_msg();
}
inline void Message::set_msg(const std::string& value) {
  _internal_set_msg(value);
  // @@protoc_insertion_point(field_set:protos.Message.msg)
}
inline std::string* Message::mutable_msg() {
  // @@protoc_insertion_point(field_mutable:protos.Message.msg)
  return _internal_mutable_msg();
}
inline const std::string& Message::_internal_msg() const {
  return msg_.Get();
}
inline void Message::_internal_set_msg(const std::string& value) {
  
  msg_.Set(::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr::EmptyDefault{}, value, GetArena());
}
inline void Message::set_msg(std::string&& value) {
  
  msg_.Set(
    ::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr::EmptyDefault{}, ::std::move(value), GetArena());
  // @@protoc_insertion_point(field_set_rvalue:protos.Message.msg)
}
inline void Message::set_msg(const char* value) {
  GOOGLE_DCHECK(value != nullptr);
  
  msg_.Set(::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr::EmptyDefault{}, ::std::string(value), GetArena());
  // @@protoc_insertion_point(field_set_char:protos.Message.msg)
}
inline void Message::set_msg(const void* value,
    size_t size) {
  
  msg_.Set(::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr::EmptyDefault{}, ::std::string(
      reinterpret_cast<const char*>(value), size), GetArena());
  // @@protoc_insertion_point(field_set_pointer:protos.Message.msg)
}
inline std::string* Message::_internal_mutable_msg() {
  
  return msg_.Mutable(::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr::EmptyDefault{}, GetArena());
}
inline std::string* Message::release_msg() {
  // @@protoc_insertion_point(field_release:protos.Message.msg)
  return msg_.Release(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), GetArena());
}
inline void Message::set_allocated_msg(std::string* msg) {
  if (msg != nullptr) {
    
  } else {
    
  }
  msg_.SetAllocated(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), msg,
      GetArena());
  // @@protoc_insertion_point(field_set_allocated:protos.Message.msg)
}

// -------------------------------------------------------------------

// Messages

// int32 status = 1;
inline void Messages::clear_status() {
  status_ = 0;
}
inline ::PROTOBUF_NAMESPACE_ID::int32 Messages::_internal_status() const {
  return status_;
}
inline ::PROTOBUF_NAMESPACE_ID::int32 Messages::status() const {
  // @@protoc_insertion_point(field_get:protos.Messages.status)
  return _internal_status();
}
inline void Messages::_internal_set_status(::PROTOBUF_NAMESPACE_ID::int32 value) {
  
  status_ = value;
}
inline void Messages::set_status(::PROTOBUF_NAMESPACE_ID::int32 value) {
  _internal_set_status(value);
  // @@protoc_insertion_point(field_set:protos.Messages.status)
}

// repeated .protos.Message msg = 2;
inline int Messages::_internal_msg_size() const {
  return msg_.size();
}
inline int Messages::msg_size() const {
  return _internal_msg_size();
}
inline void Messages::clear_msg() {
  msg_.Clear();
}
inline ::protos::Message* Messages::mutable_msg(int index) {
  // @@protoc_insertion_point(field_mutable:protos.Messages.msg)
  return msg_.Mutable(index);
}
inline ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::protos::Message >*
Messages::mutable_msg() {
  // @@protoc_insertion_point(field_mutable_list:protos.Messages.msg)
  return &msg_;
}
inline const ::protos::Message& Messages::_internal_msg(int index) const {
  return msg_.Get(index);
}
inline const ::protos::Message& Messages::msg(int index) const {
  // @@protoc_insertion_point(field_get:protos.Messages.msg)
  return _internal_msg(index);
}
inline ::protos::Message* Messages::_internal_add_msg() {
  return msg_.Add();
}
inline ::protos::Message* Messages::add_msg() {
  // @@protoc_insertion_point(field_add:protos.Messages.msg)
  return _internal_add_msg();
}
inline const ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::protos::Message >&
Messages::msg() const {
  // @@protoc_insertion_point(field_list:protos.Messages.msg)
  return msg_;
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__
// -------------------------------------------------------------------


// @@protoc_insertion_point(namespace_scope)

}  // namespace protos

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // GOOGLE_PROTOBUF_INCLUDED_GOOGLE_PROTOBUF_INCLUDED_Message_2eproto